class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        st = []
        for i in s:
            if i in ('(', '[', '{'):
                st.append(i)
            else:
                if not st:
                    return False
                last = st.pop()
                if i == ')' and last != '(':
                    return False
                if i==']' and last !='[':
                    return False
                if i=='}' and last != '{':
                    return False
        if st:
            return False
        return True
