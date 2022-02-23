package com.shu.rabbitmq;

import com.shu.entity.RankingInfo;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.annotation.RabbitHandler;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Map;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
@Component
//@RabbitListener(queues = "RankingListQueue")
@Slf4j
public class UpdateReceiver {

    private static final String RANKING_LIST = "ranking_list";

    @Resource
    RedisTemplate redisTemplate;

//    @RabbitHandler
    public void onUpdate(Map<String, Object> msg) throws ParseException {
        RankingInfo info = new RankingInfo();
        log.info("收到新的游戏成绩 {}", msg.toString());
        info.setPlayerName(msg.get("playerName").toString());
        info.setScore((Double) msg.get("score"));
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Date date = format.parse((String) msg.get("playTime"));
        info.setPlayTime(date);
        String dateStr = format.format(date);
        redisTemplate.opsForZSet().add(RANKING_LIST, info.getPlayerName() + "@" + dateStr, info.getScore());
    }
}
