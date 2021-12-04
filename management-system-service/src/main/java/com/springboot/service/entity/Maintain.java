package com.springboot.service.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "tb_maintain")
public class Maintain {
    /**
     * id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    /**
     * 用户id
     */
    private Integer userId;
    /**
     * 设备id
     */
    private Integer equipmentId;
    /**
     * 维护日期
     */
    private long date;
    /**
     * 结果
     */
    private String status;
    /**
     * 异常情况
     */
    private String remark;
}
