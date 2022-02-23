package com.shu.utils;

import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.RedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;
import org.springframework.stereotype.Component;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
@Component
public class RedisUtil {
    private static volatile RedisTemplate redisTemplate;

    private RedisUtil() {}

    public static RedisTemplate getInstance() {
        if (redisTemplate == null) {
            synchronized (RedisUtil.class) {
                if (redisTemplate == null) {
                    redisTemplate = new RedisTemplate();
                    RedisSerializer stringSerializer = new StringRedisSerializer();
                    redisTemplate.setKeySerializer(stringSerializer);
                    redisTemplate.setValueSerializer(stringSerializer);
                }
            }
        }
        return redisTemplate;
    }
}
