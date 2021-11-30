package com.springboot.service.repository;

import com.springboot.service.entity.EquipmentType;
import org.springframework.data.jpa.repository.JpaRepository;

public interface EquipmentTypeRepository extends JpaRepository<EquipmentType,Integer> {
}
