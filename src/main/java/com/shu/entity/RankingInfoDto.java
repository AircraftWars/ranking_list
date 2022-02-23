package com.shu.entity;

import lombok.Data;

import java.util.Date;

/**
 * @author oxotn3
 * @create 2022-02-23
 * @description
 */
@Data
public class RankingInfoDto {
    private String playerName;

    private double score;

    private String playTime;
}
