package com.springboot.service.repository;

import com.springboot.service.entity.Groups;
import org.springframework.data.jpa.repository.JpaRepository;

public interface GroupsRepository extends JpaRepository<Groups,Integer> {
    Groups findAllByName(String name);
}
