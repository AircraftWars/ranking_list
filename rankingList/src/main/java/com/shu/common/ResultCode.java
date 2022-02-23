package com.shu.common;

/**
 * @author oxotn3
 * @create 2022-02-22
 * @description
 */
public enum ResultCode {
    //    成功状态码
    SUCCESS(1000, "成功"),
    //    参数错误：1001-1999
    PARAM_IS_INVALID(1001, "参数无效"),
    PARAM_ISI_BLANK(1002, "参数为空"),
    PARAM_TYPE_BIND_ERROR(1003, "参数类型错误"),
    PARAM_NOT_COMPLETE(1004, "参数缺失"),

    // 未知错误
    UNKNOWN_ERROR(3000, "未知错误");


    private Integer code;
    private String message;

    ResultCode(Integer code, String message) {
        this.code = code;
        this.message = message;
    }

    public Integer getCode() {
        return this.code;
    }

    public String getMessage() {
        return this.message;
    }

}
