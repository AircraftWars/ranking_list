package com.shu.service.impl;

import com.shu.entity.RankingInfo;
import com.shu.service.RankingService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.ZSetOperations;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.*;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
@Service
@Slf4j
public class RankingServiceImpl implements RankingService {

    private static final String RANKING_LIST = "ranking_list";

    @Resource
    RedisTemplate redisTemplate;

    @Resource
    RabbitTemplate rabbitTemplate;

    @Override
    public List<RankingInfo> getTopNRankingInfo(int n) throws ParseException {
        List<RankingInfo> res = new ArrayList<>();
        Set<ZSetOperations.TypedTuple<String>> list = redisTemplate.opsForZSet().reverseRangeWithScores(RANKING_LIST, 0, n - 1);
        for (ZSetOperations.TypedTuple<String> o : list) {
            RankingInfo cur = new RankingInfo();
            cur.setScore(o.getScore());
            // value形式为 playerName@{playTime}
            String[] args = o.getValue().split("@");
            cur.setPlayerName(args[0]);
            cur.setPlayTime(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").parse(args[1]));
            res.add(cur);
            log.info("user {} scores {}, at {}" , cur.getPlayerName(), cur.getScore(), cur.getPlayTime());
        }
        return res;
    }

    @Override
    public void mqDemo() {
        Map<String, Object> map = new HashMap<>();
        String playerName = "NAME";
        double score = 100d;
        String playTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
        map.put("playerName", playerName);
        map.put("score", score);
        map.put("playTime", playTime);
        rabbitTemplate.convertAndSend("RankingListExchange", "RankingListRouting", map);
        log.info("新的游戏成绩消息发送 {}", map.toString());
    }
}
