package com.springboot.service.mapper;

import com.springboot.service.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserMapper extends JpaRepository<User,Integer> {
}
