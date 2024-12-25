#!/bin/bash
./quiz-maker create user XY
./quiz-maker create quiz test "is Test?,is not a Test?"
./quiz-maker create option 1 "Yes" true
./quiz-maker create option 1 "No" false
./quiz-maker create option 2 "Yes" false
./quiz-maker create option 2 "No" true
./quiz-maker begin 1 1
./quiz-maker answer 1 1
./quiz-maker answer 1 4
./quiz-maker submit 1
./quiz-maker get score 1 1
./quiz-maker get ranking 1 1
./quiz-maker get analysis 1 1