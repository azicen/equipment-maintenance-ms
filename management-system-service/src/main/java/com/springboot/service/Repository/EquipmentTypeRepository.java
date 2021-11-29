package com.springboot.service.Repository;

import com.springboot.service.Entity.EquipmentType;
import org.springframework.data.jpa.repository.JpaRepository;

public interface EquipmentTypeRepository extends JpaRepository<EquipmentType,Integer> {
}
