package com.shu;

import com.shu.entity.RankingInfo;
import com.shu.rabbitmq.NativeConsumer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.client.RestTemplate;

import javax.annotation.Resource;
import java.text.SimpleDateFormat;

@SpringBootApplication
public class RankingListApplication {

    public static void main(String[] args) {
        SpringApplication.run(RankingListApplication.class, args);
    }

}
