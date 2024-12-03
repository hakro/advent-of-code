// I don't want to use Regex today, it's boring.
// Let's do a brutal, error prone parser of doom

#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

int check_tok_mul(FILE *f);   // is next token 'mul'
int check_tok_op(FILE *f); 	  // is next token open '('
int check_tok_cp(FILE *f);    // is next token close ')'
int check_tok_comma(FILE *f); // is next token ','
int check_tok_num(FILE *f);   // is next token number

int main() {
	FILE *f = fopen("input.txt", "r");
	if (f == NULL) {
		printf("err: cannot open input file");
		return 1;
	}
	
	// Let's do some if-else butchering, I want to go to sleep
	int final_res = 0;
	while (!feof(f)) {
		if (check_tok_mul(f)) {
			if (check_tok_op(f)) {
				int n1 = check_tok_num(f);
				if (n1) {
					if (check_tok_comma(f)) {
						int n2 = check_tok_num(f);
						if (n2) {
							if (check_tok_cp(f)) {
								/* printf("%d, %d\n", n1, n2); */
								final_res += n1 * n2;
							}
						}
					}
				}
				
			}
		}
	}
	fclose(f);

	printf("Final result: %d\n", final_res);
}

// Check if the next token is the keyword 'mul'
int check_tok_mul(FILE *f) {
	char c = fgetc(f);
	if (c != 'm') {
		return 0;
	}
	c = fgetc(f);
	if (c != 'u') {
		return 0;
	}
	c = fgetc(f);
	if (c != 'l') {
		return 0;
	}
	return 1;
}

int check_tok_op(FILE *f) {
	char c = fgetc(f);
	if (c != '(') {
		return 0;
	}
	return 1;
} 

int check_tok_cp(FILE *f) {
	char c = fgetc(f);
	if (c != ')') {
		return 0;
	}
	return 1;
} 

int check_tok_comma(FILE *f) {
	char c = fgetc(f);
	if (c != ',') {
		return 0;
	}
	return 1;
} 

int check_tok_num(FILE *f) {
	char num_buff[10];
	while (1) {
		char c = fgetc(f);
		if (!isdigit(c)) {
			// Not a digit, so go back a step, and leave
			fseek(f, -1, SEEK_CUR);
			break;
		}
		strncat(num_buff, &c, 1);
	}
	if (strlen(num_buff)== 0) {
		return 0;
	} 
	return atoi(num_buff);
}
