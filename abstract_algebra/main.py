from itertools import permutations
import time

def main():
    n = 4 # 4阶转置群

    ## 生成置换群
    permutations = generate_permutation_group(n)
    print(permutations)
    print(len(permutations))
    for permutation in permutations:
        print("||||||||||||||||||||||")
        print(permutation)
        print(permutation_order(permutation))
        print("______________________")


    ### 结论： 4阶置换群有 24 个置换，其中有 6 个阶为 4，有 8 个阶为 3，有 6 个阶为 2，有 3 个阶为 1
    ### S4 不能表示为Z3*Z8
        
## 求置换群的阶
def permutation_order(permutation):
    """
    Compute the order of a permutation.
    
    Parameters:
        permutation (tuple): A permutation of {1, 2, ..., n}.
        
    Returns:
        int: The order of the permutation.
    """
    n = len(permutation)

    mapping = {}
    for i in range(n):
        mapping[i + 1] = permutation[i]

    ## target = [1, 2, ..., n]
    target = list(range(1, n + 1))
    permutation = list(permutation)
    
    count = 1
    while True:
        print(permutation)
        if permutation == target:
            break
        nextPermutation = [mapping[k] for k in permutation]
        permutation = nextPermutation
        count += 1

    return count


def generate_permutation_group(n):
    """
    Generate the permutation group S_n.
    
    Parameters:
        n (int): The number of elements in the set {1, 2, ..., n}.
        
    Returns:
        list of tuples: All permutations of {1, 2, ..., n}.
    """
    if n <= 0:
        raise ValueError("n must be a positive integer.")
    
    elements = list(range(1, n + 1))
    return list(permutations(elements))
    



if __name__ == '__main__':
    main()