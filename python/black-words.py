# 整个烂活，灵感来源：狗屁不通文章生成器
import random

stencil = '{n40}是{v0}{n41}，{v1}行业{n30}。{n42}是{v2}{n20}{n43}，通过{n31}和{n32}达到{n33}。' \
          '{n44}是在{n45}采用{n21}打法达成{n46}。{n47}{n48}作为{n22}为产品赋能，{n49}作为{n23}' \
          '的评判标准。亮点是{n24}，优势是{n25}。{v3}整个{n410}，{v4}{n26}{v5}{n411}。{n34}是{n35}' \
          '达到{n36}标准。'

num = {'v': 6, 'n2': 7, 'n3': 7, 'n4': 12}


# 二字动词
v = '皮实、复盘、赋能、加持、沉淀、倒逼、落地、串联、协同、反哺、兼容、包装、重组、履约、' \
    '响应、量化、发力、布局、联动、细分、梳理、输出、加速、共建、共创、支撑、融合、解耦、聚合、' \
    '集成、对标、对齐、聚焦、抓手、拆解、拉通、抽象、摸索、提炼、打通、吃透、迁移、分发、分层、' \
    '封装、辐射、围绕、复用、渗透、扩展、开拓、给到、死磕、破圈'.split('、')

# 二字名词
n2 = '漏斗、中台、闭环、打法、纽带、矩阵、刺激、规模、场景、维度、格局、形态、生态、话术、' \
     '体系、认知、玩法、体感、感知、调性、心智、战役、合力、赛道、基因、因子、模型、载体、横向、' \
     '通道、补位、链路、试点'.split('、')

# 三字名词
n3 = '新生态、感知度、颗粒度、方法论、组合拳、引爆点、点线面、精细化、差异化、平台化、结构化、' \
     '影响力、耦合性、易用性、便捷性、一致性、端到端、短平快、护城河'.split('、')

# 四字名词
n4 = '底层逻辑、顶层设计、交付价值、生命周期、价值转化、强化认知、资源倾斜、完善逻辑、抽离透传、' \
     '复用打法、商业模式、快速响应、定性定量、关键路径、去中心化、结果导向、垂直领域、归因分析、' \
     '体验度量、信息屏障'.split('、')

v_list = random.sample(v, num['v'])
n2_list = random.sample(n2, num['n2'])
n3_list = random.sample(n3, num['n3'])
n4_list = random.sample(n4, num['n4'])
lists = {'v': v_list, 'n2': n2_list, 'n3': n3_list, 'n4': n4_list}

dic = {}
for current_type in ['v', 'n2', 'n3', 'n4']:
    current_list = lists[current_type]
    for i in range(0, len(current_list)):
        dic[current_type + str(i)] = current_list[i]

result = stencil.format(**dic)
print(result)

''' 
输出结果：

生命周期是发力快速响应，赋能行业引爆点。商业模式是细分载体体验度量，通过平台化和便捷性达到短平快。
完善逻辑是在底层逻辑采用玩法打法达成强化认知。复用打法资源倾斜作为打法为产品赋能，信息屏障作为体
系的评判标准。亮点是维度，优势是闭环。聚焦整个顶层设计，扩展规模迁移垂直领域。颗粒度是组合拳达到影
响力标准。

代码水平低下，仅图一乐，还请轻喷编辑于 03-31
'''