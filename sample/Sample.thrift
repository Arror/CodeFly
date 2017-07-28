namespace swift TR

struct Base {
     
}

struct Example {
    0: required i16 r_i16_v
    1: required i32 r_i32_v
    2: required i64 r_i64_v
    3: required double r_double_v
    4: required bool r_bool_v
    5: required string r_string_v
    6: required Base r_base_v
    8: required list<i16> r_i16_v_list
    9: required list<i32> r_i32_v_list
    10: required list<i64> r_i64_v_list
    11: required list<double> r_double_v_list
    12: required list<bool> r_bool_v_list
    13: required list<string> r_string_v_list
    14: required list<Base> r_base_v_list
    16: optional i16 o_i16_v
    17: optional i32 o_i32_v
    18: optional i64 o_i64_v
    19: optional double o_double_v
    20: optional bool o_bool_v
    21: optional string o_string_v
    22: optional Base o_base_v
    24: optional list<i16> o_i16_v_list
    25: optional list<i32> o_i32_v_list
    26: optional list<i64> o_i64_v_list
    27: optional list<double> o_double_v_list
    28: optional list<bool> o_bool_v_list
    29: optional list<string> o_string_v_list
    30: optional list<Base> o_base_v_list
}

struct ExampleReq {

}

struct ExampleResp {

}

service Example {

    ExampleResp ping(1: optional ExampleReq ping_req, 2: required string user_name) (PATH = "test/ping");

    ExampleResp pint() (PATH = "test/pint")
}