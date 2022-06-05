import random
from typing import List


# """
# This is Master's API interface.
# You should not implement it, or speculate about its implementation
# """
# class Master:
#     def guess(self, word: str) -> int:


class Solution:
    def findSecretWord(self, wordlist: List[str], master: 'Master') -> None:
        for i in range(10):
            current_word = wordlist[random.randrange(len(wordlist))]
            matches = master.guess(current_word)
            if matches == 6:
                break
            next_wordlist = []
            for word in wordlist:
                word_matches = 0
                for j in range(6):
                    if current_word[j] == word[j]:
                        word_matches += 1
                if word_matches == matches:
                    next_wordlist.append(word)
            wordlist = next_wordlist
