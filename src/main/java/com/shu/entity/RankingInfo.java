package com.shu.entity;

import lombok.Data;

import java.io.Serializable;
import java.util.Date;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
@Data
public class RankingInfo implements Serializable {
    private String playerName;

    private double score;

    private Date playTime;
}
