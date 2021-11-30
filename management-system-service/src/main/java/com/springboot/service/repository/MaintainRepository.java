package com.springboot.service.repository;

import com.springboot.service.entity.Maintain;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MaintainRepository extends JpaRepository<Maintain,Integer> {
}
