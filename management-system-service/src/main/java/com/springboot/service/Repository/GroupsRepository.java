package com.springboot.service.Repository;

import com.springboot.service.Entity.Groups;
import org.springframework.data.jpa.repository.JpaRepository;

public interface GroupsRepository extends JpaRepository<Groups,Integer> {
}
