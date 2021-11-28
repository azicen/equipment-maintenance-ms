package com.springboot.service.Repository;

import com.springboot.service.Entity.Machine;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MachineRepository extends JpaRepository<Machine,Integer> {
}
