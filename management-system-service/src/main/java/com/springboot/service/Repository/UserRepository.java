package com.springboot.service.Repository;

import com.springboot.service.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface UserRepository extends JpaRepository<User,Integer> {
    List<User> findByName(String name);
    Integer removeAllById(Integer id);
}
