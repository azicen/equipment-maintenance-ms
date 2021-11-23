package com.springboot.server;

import com.springboot.server.entity.User;
import com.springboot.server.mapper.UserMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;
import java.util.List;

@SpringBootTest
class ManagementSystemServiceApplicationTests {

    @Resource
    private UserMapper userMapper;

    @Test
    public void testSelect() {
        User user=new User();
        user.setName("null13");
        user.setAge(18);
//        user.setId(9l);
        user.setEmail("1229608477@qq.com");
        int insert = userMapper.insert(user);
        System.out.println(insert);
    }

}
