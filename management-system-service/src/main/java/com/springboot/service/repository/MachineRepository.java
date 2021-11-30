package com.springboot.service.repository;

import com.springboot.service.entity.Equipment;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MachineRepository extends JpaRepository<Equipment,Integer> {
}
