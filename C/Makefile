# 컴파일러와 플래그 정의
CC = gcc
CFLAGS = -Wall -O3

# 소스 파일과 타겟 정의
SRC = main.c insertion/insertion.c
TARGET = output

# 빌드 규칙
all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC)

# clean 규칙
clean:
	rm -f $(TARGET)

# 객체 파일을 생성하고 싶을 경우
# OBJ = main.o insertion/insertion.o
# $(TARGET): $(OBJ)
#	$(CC) $(CFLAGS) -o $(TARGET) $(OBJ)
