package com.shu.service;

import com.shu.entity.RankingInfo;

import java.text.ParseException;
import java.util.List;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
public interface RankingService {
    List<RankingInfo> getTopNRankingInfo(int n) throws ParseException;
    void mqDemo();
}
