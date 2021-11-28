package com.springboot.service.Repository;

import com.springboot.service.Entity.Maintain;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MaintainRepository extends JpaRepository<Maintain,Integer> {
}
