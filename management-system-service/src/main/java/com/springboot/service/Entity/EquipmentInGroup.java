package com.springboot.service.Entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "tb_equipment_type")
public class EquipmentInGroup {
    /**
     * id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    /**
     * 权限组编号
     */
    private Integer groupId;
    /**
     * 设备编号
     */
    private Integer equipmentId;
}
