package com.shu.rabbitmq;

import com.rabbitmq.client.*;
import com.shu.entity.RankingInfo;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Map;
import java.util.concurrent.TimeUnit;

/**
 * @author oxotn3
 * @create 2022-02-23
 * @description
 */
@Component
@Slf4j
public class NativeConsumer {
    private static final String QUEUE_NAME = "RankingListQueue";
    private static final String IP_ADDRESS = "127.0.0.1";
    private static final int PORT = 5672;

    private static final String RANKING_LIST = "ranking_list";

    @Resource
    private RedisTemplate redisTemplate;


    @PostConstruct
    public void func() throws Exception{
        Address[] addresses = new Address[] { new Address(IP_ADDRESS, PORT) };
        ConnectionFactory factory = new ConnectionFactory();
        factory.setUsername("guest");
        factory.setPassword("guest");
        // 这里的连接方式与生产者的demo略有不同，注意辨别区别
        // 创建连接
        Connection connection = factory.newConnection(addresses);
        // 创建信道
        final Channel channel = connection.createChannel();
        // 设置客户端最多接受未被ack的消息的个数
        channel.basicQos(64);
        Consumer consumer = new DefaultConsumer(channel) {
            @SneakyThrows
            @Override
            public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                ByteArrayInputStream bi = new ByteArrayInputStream(body);
                ObjectInputStream oi = new ObjectInputStream(bi);
                Object obj = null;
                try {
                    obj = oi.readObject();
                } catch (ClassNotFoundException e) {
                    e.printStackTrace();
                }
                bi.close();
                oi.close();
                Map<String, Object> msg = (Map<String, Object>) obj;
                log.info("收到新的游戏成绩 {}", msg.toString());

//                UpdateReceiver receiver = new UpdateReceiver();
//                receiver.onUpdate(msg);

                RankingInfo info = new RankingInfo();
                info.setPlayerName(msg.get("playerName").toString());
                info.setScore((Double) msg.get("score"));
                SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
                Date date = null;
                try {
                    date = format.parse((String) msg.get("playTime"));
                } catch (ParseException e) {
                    e.printStackTrace();
                }
                info.setPlayTime(date);
                updateToRedis(info);

                try {
                    TimeUnit.SECONDS.sleep(1);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                channel.basicAck(envelope.getDeliveryTag(), false);
            }
        };
        channel.basicConsume(QUEUE_NAME, consumer);
        // 等待回调函数执行完毕之后，关闭资源
        TimeUnit.SECONDS.sleep(5);
//        channel.close();
//        connection.close();
    }

    private void updateToRedis(RankingInfo info) {
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        String dateStr = format.format(info.getPlayTime());
//        log.info("{}==={}==={}", info.getPlayerName(), info.getScore(), info.getPlayTime());
        redisTemplate.opsForZSet().add(RANKING_LIST, info.getPlayerName() + "@" + dateStr, info.getScore());
    }
}

