package com.springboot.service.Entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.sql.Timestamp;

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
    private Integer machineId;
    /**
     * 维护日期
     */
    private Timestamp date;
    /**
     * 结果
     */
    private String status;
    /**
     * 异常情况
     */
    private String remark;
}
