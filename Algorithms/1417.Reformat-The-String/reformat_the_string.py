class Solution:
    def reformat(self, s: str) -> str:
        length = len(s)
        arr = ['' for i in range(length + 1)]
        digit_pos = 1
        str_pos = 2
        for ele in s:
            if str.isdigit(ele):
                if digit_pos > length:
                    return ""
                arr[digit_pos] = ele
                digit_pos += 2
            else:
                if str_pos == length + 1:
                    arr[0] = ele
                    str_pos += 2
                    continue
                if str_pos > length + 1:
                    return ""
                arr[str_pos] = ele
                str_pos += 2
        return "".join(arr)
        
