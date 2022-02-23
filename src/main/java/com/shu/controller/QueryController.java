package com.shu.controller;

import com.shu.common.Result;
import com.shu.common.ResultCode;
import com.shu.entity.RankingInfo;
import com.shu.entity.RankingInfoDto;
import com.shu.service.RankingService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.text.ParseException;
import java.util.List;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description 查询排名
 */
@Slf4j
@RestController
@RequestMapping("/api")
public class QueryController {
    @Resource
    RankingService rankingService;

    @GetMapping("/getTopNRankingInfo")
    @ResponseBody
    public Result getTopNRankingInfo(@RequestParam int n) {
        if (n <= 0) return Result.failure(ResultCode.PARAM_IS_INVALID);
        List<RankingInfoDto> infoList;
        try {
            infoList = rankingService.getTopNRankingInfo(n);
        } catch (ParseException e) {
            e.printStackTrace();
            return Result.failure(ResultCode.UNKNOWN_ERROR);
        }
        return Result.success(infoList);
    }

    @GetMapping("/mq")
    @ResponseBody
    public String mq() {
        rankingService.mqDemo();
        return "mq";
    }
}
