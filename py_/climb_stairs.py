# -*- coding:utf-8 _*-
"""
@author:lee 
@file: climb_stairs.py 
@time: 2021/05/20
@contact: leeguo2015@163.com
@site:  
@software: PyCharm 

"""

def climb_stairs( step):
    if step == 1:
        return 1
    if step == 2:
        return 2
    return climb_stairs(step-1) +climb_stairs(step-2)

if __name__ == "__main__":
    result = climb_stairs(4)
    print(result)
