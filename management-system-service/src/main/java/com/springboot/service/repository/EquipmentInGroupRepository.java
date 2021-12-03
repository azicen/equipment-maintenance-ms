package com.springboot.service.repository;

import com.springboot.service.entity.EquipmentInGroup;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface EquipmentInGroupRepository extends JpaRepository<EquipmentInGroup,Integer> {
    EquipmentInGroup findAllByEquipmentIdAndGroupId(Integer eId,Integer groupId);
    List<EquipmentInGroup> findAllByEquipmentId(Integer id);
}
