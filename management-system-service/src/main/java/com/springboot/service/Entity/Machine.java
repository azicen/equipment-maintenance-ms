package com.springboot.service.Entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.sql.Date;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "tb_machine")
@JsonIgnoreProperties(value={"hibernateLazyInitializer","handler","fieldHandler"})
public class Machine {
    /**
     * 设备id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    /**
     * 设备姓名
     */
    private String name;
    /**
     * 地点
     */
    private String location;
    /**
     * 状况
     */
    private String status;
    /**
     * 近期维护日期
     */
    private Date date;
    /**
     * 开始服役日期
     */
    private Date startDate;
    /**
     * 结束服务日期
     */
    private Date deadLine;
    /**
     * 类型编号
     */
    private String groupId;
    /**
     * 创建人员编号
     */
    private Integer userId;
    /**
     * 创建日期
     */
    private Date createDate;
}
