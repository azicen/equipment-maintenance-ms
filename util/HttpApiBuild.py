import os


def change_variable_name(listx):
    result = ''
    for i in range(len(listx)):
        if listx[i].isupper():
            if i > 0:
                if not listx[i - 1].isupper():
                    result += "_"
            result += listx[i].lower()
        else:
            result += listx[i]
    return result


def get_property(str):
    s_arr = str.split("(")
    model_name = s_arr[0].strip()

    s_arr = s_arr[1:]
    for i in range(len(s_arr)):
        s_arr[i] = s_arr[i].strip()
        if(s_arr[i][-1] == ")"):
            s_arr[i] = s_arr[i][0:-1]
            if s_arr[i] == "":
                s_arr[i] = "%null%"

    p_dict = {}
    r_dict = {}
    for i in range(len(s_arr)):
        ss_arr = s_arr[i].split(",")
        if i < (len(s_arr) / 2):
            for j in range(len(ss_arr)):
                ss_arr[j] = ss_arr[j].strip()
                if ss_arr[j] == "%null%":
                    continue
                p_arr = ss_arr[j].split(" ")
                p_name = p_arr[0].strip()
                p_type = p_arr[1].strip()
                p_dict[p_name] = p_type
        else:
            for j in range(len(ss_arr)):
                ss_arr[j] = ss_arr[j].strip()
                if ss_arr[j] == "%null%":
                    continue
                r_arr = ss_arr[j].split(" ")
                r_name = r_arr[0].strip()
                r_type = r_arr[1].strip()
                r_dict[r_name] = r_type
    return model_name, p_dict, r_dict


def build_model(name, v_dict, n, fo):
    fo.write("type HTTP" + name + n + " struct {\n")
    for (p_name, p_type) in v_dict.items():
        n = change_variable_name(p_name)
        fo.write("    " + p_name + " " + p_type +
                 " `form:\"" + n + "\" json:\"" + n + "\" binding:\"required\"`\n")
    fo.write("}\n")
    fo.write("\n")


def build_info_model(name, p_dict, fo):
    fo.write("//HTTP" + name + "Info HTTP消息模型" + name + "\n")
    build_model(name, p_dict, "Info", fo)


def build_response_model(name, r_dict, fo):
    fo.write("//HTTP" + name + "Info HTTP回复模型" + name + "\n")
    build_model(name, r_dict, "Response", fo)


def build_dao_func(name, p_dict, fo):
    model_name = "HTTP" + name + "Info"
    fo.write("//Bind" + model_name + " \n")
    fo.write("func (d *Dao) Bind" + model_name +
             "(c *core.Context) (info model." + model_name + ", err error) {\n")

    fo.write("    err = c.BindJSON(&info)\n")
    fo.write("    if err != nil {\n        return\n    }\n")

    fo.write("    return\n")
    fo.write("}\n")
    fo.write("\n")


def build_api_func(name, r_dict, fo):
    model_name = "HTTP" + name + "Response"
    fo.write("//" + name + " " + name + "服务api逻辑处理\n")
    fo.write("func (s *Service) " + name + "(c *core.Context) {\n")

    fo.write("    info, err := s.dao.BindHTTP" + name + "Info(c)\n")
    fo.write("    if err != nil {\n")
    fo.write("        log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)\n")
    fo.write("        return\n    }\n")

    fo.write("\n\n\n")

    fo.write("    response := model." + model_name + "{\n")
    for r_name in r_dict:
        fo.write("        " + r_name + ":  ,\n")
    fo.write("    }\n")
    fo.write("    c.ReturnSuccess(response)\n")

    fo.write("}\n")
    fo.write("\n")


def build(model_name, p_dict, r_dict):
    fo = open("build/" + model_name + ".txt", "a+")
    build_info_model(model_name, p_dict, fo)
    build_response_model(model_name, r_dict, fo)
    build_dao_func(model_name, p_dict, fo)
    build_api_func(model_name, r_dict, fo)
    fo.close()


def run():
    f = open("api.txt", encoding='utf-8')
    isExists = os.path.exists("build")
    if not isExists:
        os.makedirs("build")

    while True:
        line = f.readline().strip()
        if line:
            if line[0] != "#":
                model_name, p_dict, r_dict = get_property(line)
                print("模型名称: " + model_name)
                build(model_name, p_dict, r_dict)
        else:
            break
    f.close()

    # HTTPModelName(p1 type, p2 type...) (r1 type, r2 type)
    # str = input()
    # str = "ModelName(p1 type, p2 type) (r1 type, r2 type)"


if __name__ == '__main__':
    run()
