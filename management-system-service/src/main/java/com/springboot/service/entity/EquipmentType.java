package com.springboot.service.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "tb_equipment_type")
public class EquipmentType {
    /**
     * id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    /**
     * 设备类型姓名
     */
    private String name;
    /**
     * 维护周期
     */
    private Integer cycle;
    /**
     * 权限组编号
     */
    private Integer groupId;
}
