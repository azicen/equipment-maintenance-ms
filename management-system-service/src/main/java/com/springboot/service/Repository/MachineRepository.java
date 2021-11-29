package com.springboot.service.Repository;

import com.springboot.service.Entity.Equipment;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MachineRepository extends JpaRepository<Equipment,Integer> {
}
