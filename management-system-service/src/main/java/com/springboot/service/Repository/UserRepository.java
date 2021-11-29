package com.springboot.service.Repository;

import com.springboot.service.Entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.transaction.annotation.Transactional;


public interface UserRepository extends JpaRepository<User,Integer> {
    @Transactional
    void deleteAllById(Integer id);
    User findAllByIdAndPasswd(Integer id,String passwd);
}
