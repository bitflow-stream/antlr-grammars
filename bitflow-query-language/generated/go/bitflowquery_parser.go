// Code generated from BitflowQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // BitflowQuery
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 354,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 5, 2, 87, 10,
	2, 3, 2, 5, 2, 90, 10, 2, 3, 2, 5, 2, 93, 10, 2, 3, 2, 5, 2, 96, 10, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 106, 10, 3, 12, 3,
	14, 3, 109, 11, 3, 3, 3, 5, 3, 112, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 131, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 142, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 148, 10, 7, 12, 7, 14,
	7, 151, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 163, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 187, 10, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 7, 17, 193, 10, 17, 12, 17, 14, 17, 196, 11, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 207, 10, 19, 12,
	19, 14, 19, 210, 11, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 5, 23, 224, 10, 23, 3, 24, 3, 24, 3,
	25, 5, 25, 229, 10, 25, 3, 25, 5, 25, 232, 10, 25, 3, 25, 5, 25, 235, 10,
	25, 3, 25, 5, 25, 238, 10, 25, 3, 25, 5, 25, 241, 10, 25, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 251, 10, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 257, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30,
	263, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 269, 10, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 5, 32, 284, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 7, 32, 294, 10, 32, 12, 32, 14, 32, 297, 11, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 307, 10, 33, 12, 33, 14,
	33, 310, 11, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	7, 34, 320, 10, 34, 12, 34, 14, 34, 323, 11, 34, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 7, 42, 347, 10,
	42, 12, 42, 14, 42, 350, 11, 42, 3, 42, 3, 42, 3, 42, 2, 4, 12, 62, 43,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 80, 82, 2, 9, 4, 2, 23, 23, 36, 36, 3, 2, 50, 52, 3, 2, 50, 53,
	3, 2, 45, 49, 3, 2, 40, 41, 3, 2, 34, 37, 3, 2, 43, 44, 2, 356, 2, 84,
	3, 2, 2, 2, 4, 111, 3, 2, 2, 2, 6, 113, 3, 2, 2, 2, 8, 130, 3, 2, 2, 2,
	10, 132, 3, 2, 2, 2, 12, 141, 3, 2, 2, 2, 14, 152, 3, 2, 2, 2, 16, 154,
	3, 2, 2, 2, 18, 162, 3, 2, 2, 2, 20, 164, 3, 2, 2, 2, 22, 168, 3, 2, 2,
	2, 24, 172, 3, 2, 2, 2, 26, 176, 3, 2, 2, 2, 28, 180, 3, 2, 2, 2, 30, 186,
	3, 2, 2, 2, 32, 188, 3, 2, 2, 2, 34, 199, 3, 2, 2, 2, 36, 202, 3, 2, 2,
	2, 38, 211, 3, 2, 2, 2, 40, 214, 3, 2, 2, 2, 42, 217, 3, 2, 2, 2, 44, 223,
	3, 2, 2, 2, 46, 225, 3, 2, 2, 2, 48, 228, 3, 2, 2, 2, 50, 242, 3, 2, 2,
	2, 52, 244, 3, 2, 2, 2, 54, 250, 3, 2, 2, 2, 56, 256, 3, 2, 2, 2, 58, 262,
	3, 2, 2, 2, 60, 268, 3, 2, 2, 2, 62, 283, 3, 2, 2, 2, 64, 298, 3, 2, 2,
	2, 66, 313, 3, 2, 2, 2, 68, 326, 3, 2, 2, 2, 70, 332, 3, 2, 2, 2, 72, 334,
	3, 2, 2, 2, 74, 336, 3, 2, 2, 2, 76, 338, 3, 2, 2, 2, 78, 340, 3, 2, 2,
	2, 80, 342, 3, 2, 2, 2, 82, 348, 3, 2, 2, 2, 84, 86, 5, 4, 3, 2, 85, 87,
	5, 38, 20, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2,
	2, 88, 90, 5, 36, 19, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92,
	3, 2, 2, 2, 91, 93, 5, 42, 22, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 95, 3, 2, 2, 2, 94, 96, 5, 40, 21, 2, 95, 94, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 7, 2, 2, 3, 98, 3, 3, 2, 2, 2,
	99, 100, 7, 18, 2, 2, 100, 112, 5, 6, 4, 2, 101, 107, 7, 18, 2, 2, 102,
	103, 5, 8, 5, 2, 103, 104, 7, 3, 2, 2, 104, 106, 3, 2, 2, 2, 105, 102,
	3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2,
	2, 2, 108, 110, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 112, 5, 8, 5, 2,
	111, 99, 3, 2, 2, 2, 111, 101, 3, 2, 2, 2, 112, 5, 3, 2, 2, 2, 113, 114,
	9, 2, 2, 2, 114, 7, 3, 2, 2, 2, 115, 116, 5, 12, 7, 2, 116, 117, 7, 24,
	2, 2, 117, 118, 7, 51, 2, 2, 118, 131, 3, 2, 2, 2, 119, 131, 5, 12, 7,
	2, 120, 121, 5, 18, 10, 2, 121, 122, 7, 24, 2, 2, 122, 123, 7, 51, 2, 2,
	123, 131, 3, 2, 2, 2, 124, 131, 5, 18, 10, 2, 125, 126, 5, 10, 6, 2, 126,
	127, 7, 24, 2, 2, 127, 128, 7, 51, 2, 2, 128, 131, 3, 2, 2, 2, 129, 131,
	5, 10, 6, 2, 130, 115, 3, 2, 2, 2, 130, 119, 3, 2, 2, 2, 130, 120, 3, 2,
	2, 2, 130, 124, 3, 2, 2, 2, 130, 125, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2,
	131, 9, 3, 2, 2, 2, 132, 133, 9, 3, 2, 2, 133, 11, 3, 2, 2, 2, 134, 135,
	8, 7, 1, 2, 135, 136, 5, 14, 8, 2, 136, 137, 5, 12, 7, 2, 137, 138, 5,
	16, 9, 2, 138, 142, 3, 2, 2, 2, 139, 142, 5, 18, 10, 2, 140, 142, 5, 10,
	6, 2, 141, 134, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2, 2, 2,
	142, 149, 3, 2, 2, 2, 143, 144, 12, 5, 2, 2, 144, 145, 5, 78, 40, 2, 145,
	146, 5, 12, 7, 6, 146, 148, 3, 2, 2, 2, 147, 143, 3, 2, 2, 2, 148, 151,
	3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 13, 3, 2,
	2, 2, 151, 149, 3, 2, 2, 2, 152, 153, 7, 38, 2, 2, 153, 15, 3, 2, 2, 2,
	154, 155, 7, 39, 2, 2, 155, 17, 3, 2, 2, 2, 156, 163, 5, 20, 11, 2, 157,
	163, 5, 22, 12, 2, 158, 163, 5, 24, 13, 2, 159, 163, 5, 26, 14, 2, 160,
	163, 5, 28, 15, 2, 161, 163, 5, 30, 16, 2, 162, 156, 3, 2, 2, 2, 162, 157,
	3, 2, 2, 2, 162, 158, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 162, 160, 3, 2,
	2, 2, 162, 161, 3, 2, 2, 2, 163, 19, 3, 2, 2, 2, 164, 165, 7, 25, 2, 2,
	165, 166, 7, 51, 2, 2, 166, 167, 7, 39, 2, 2, 167, 21, 3, 2, 2, 2, 168,
	169, 7, 26, 2, 2, 169, 170, 7, 51, 2, 2, 170, 171, 7, 39, 2, 2, 171, 23,
	3, 2, 2, 2, 172, 173, 7, 27, 2, 2, 173, 174, 7, 51, 2, 2, 174, 175, 7,
	39, 2, 2, 175, 25, 3, 2, 2, 2, 176, 177, 7, 28, 2, 2, 177, 178, 7, 51,
	2, 2, 178, 179, 7, 39, 2, 2, 179, 27, 3, 2, 2, 2, 180, 181, 7, 29, 2, 2,
	181, 182, 7, 51, 2, 2, 182, 183, 7, 39, 2, 2, 183, 29, 3, 2, 2, 2, 184,
	187, 5, 32, 17, 2, 185, 187, 5, 34, 18, 2, 186, 184, 3, 2, 2, 2, 186, 185,
	3, 2, 2, 2, 187, 31, 3, 2, 2, 2, 188, 189, 7, 30, 2, 2, 189, 194, 7, 51,
	2, 2, 190, 191, 7, 4, 2, 2, 191, 193, 7, 51, 2, 2, 192, 190, 3, 2, 2, 2,
	193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	197, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 198, 7, 39, 2, 2, 198, 33,
	3, 2, 2, 2, 199, 200, 7, 30, 2, 2, 200, 201, 7, 5, 2, 2, 201, 35, 3, 2,
	2, 2, 202, 203, 7, 21, 2, 2, 203, 208, 5, 52, 27, 2, 204, 205, 7, 3, 2,
	2, 205, 207, 5, 52, 27, 2, 206, 204, 3, 2, 2, 2, 207, 210, 3, 2, 2, 2,
	208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 37, 3, 2, 2, 2, 210, 208,
	3, 2, 2, 2, 211, 212, 7, 19, 2, 2, 212, 213, 5, 62, 32, 2, 213, 39, 3,
	2, 2, 2, 214, 215, 7, 22, 2, 2, 215, 216, 5, 62, 32, 2, 216, 41, 3, 2,
	2, 2, 217, 218, 7, 20, 2, 2, 218, 219, 5, 44, 23, 2, 219, 43, 3, 2, 2,
	2, 220, 224, 5, 46, 24, 2, 221, 224, 5, 50, 26, 2, 222, 224, 5, 48, 25,
	2, 223, 220, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224,
	45, 3, 2, 2, 2, 225, 226, 7, 23, 2, 2, 226, 47, 3, 2, 2, 2, 227, 229, 7,
	32, 2, 2, 228, 227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 231, 3, 2, 2,
	2, 230, 232, 5, 54, 28, 2, 231, 230, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2,
	232, 234, 3, 2, 2, 2, 233, 235, 5, 56, 29, 2, 234, 233, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 238, 5, 58, 30, 2, 237, 236,
	3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 240, 3, 2, 2, 2, 239, 241, 5, 60,
	31, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 49, 3, 2, 2, 2,
	242, 243, 7, 50, 2, 2, 243, 51, 3, 2, 2, 2, 244, 245, 7, 51, 2, 2, 245,
	53, 3, 2, 2, 2, 246, 247, 7, 50, 2, 2, 247, 251, 7, 6, 2, 2, 248, 249,
	7, 50, 2, 2, 249, 251, 7, 7, 2, 2, 250, 246, 3, 2, 2, 2, 250, 248, 3, 2,
	2, 2, 251, 55, 3, 2, 2, 2, 252, 253, 7, 50, 2, 2, 253, 257, 7, 8, 2, 2,
	254, 255, 7, 50, 2, 2, 255, 257, 7, 9, 2, 2, 256, 252, 3, 2, 2, 2, 256,
	254, 3, 2, 2, 2, 257, 57, 3, 2, 2, 2, 258, 259, 7, 50, 2, 2, 259, 263,
	7, 10, 2, 2, 260, 261, 7, 50, 2, 2, 261, 263, 7, 11, 2, 2, 262, 258, 3,
	2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 59, 3, 2, 2, 2, 264, 265, 7, 50, 2,
	2, 265, 269, 7, 12, 2, 2, 266, 267, 7, 50, 2, 2, 267, 269, 7, 13, 2, 2,
	268, 264, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 61, 3, 2, 2, 2, 270, 271,
	8, 32, 1, 2, 271, 272, 7, 38, 2, 2, 272, 273, 5, 62, 32, 2, 273, 274, 7,
	39, 2, 2, 274, 284, 3, 2, 2, 2, 275, 276, 5, 70, 36, 2, 276, 277, 5, 62,
	32, 10, 277, 284, 3, 2, 2, 2, 278, 284, 5, 80, 41, 2, 279, 284, 5, 66,
	34, 2, 280, 284, 5, 64, 33, 2, 281, 284, 5, 68, 35, 2, 282, 284, 5, 72,
	37, 2, 283, 270, 3, 2, 2, 2, 283, 275, 3, 2, 2, 2, 283, 278, 3, 2, 2, 2,
	283, 279, 3, 2, 2, 2, 283, 280, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283,
	282, 3, 2, 2, 2, 284, 295, 3, 2, 2, 2, 285, 286, 12, 9, 2, 2, 286, 287,
	5, 74, 38, 2, 287, 288, 5, 62, 32, 10, 288, 294, 3, 2, 2, 2, 289, 290,
	12, 8, 2, 2, 290, 291, 5, 76, 39, 2, 291, 292, 5, 62, 32, 9, 292, 294,
	3, 2, 2, 2, 293, 285, 3, 2, 2, 2, 293, 289, 3, 2, 2, 2, 294, 297, 3, 2,
	2, 2, 295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 63, 3, 2, 2, 2,
	297, 295, 3, 2, 2, 2, 298, 299, 7, 31, 2, 2, 299, 300, 7, 51, 2, 2, 300,
	301, 7, 39, 2, 2, 301, 302, 7, 33, 2, 2, 302, 303, 7, 14, 2, 2, 303, 308,
	7, 51, 2, 2, 304, 305, 7, 4, 2, 2, 305, 307, 7, 51, 2, 2, 306, 304, 3,
	2, 2, 2, 307, 310, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2,
	2, 309, 311, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 311, 312, 7, 15, 2, 2, 312,
	65, 3, 2, 2, 2, 313, 314, 7, 51, 2, 2, 314, 315, 7, 33, 2, 2, 315, 316,
	7, 14, 2, 2, 316, 321, 7, 50, 2, 2, 317, 318, 7, 4, 2, 2, 318, 320, 7,
	50, 2, 2, 319, 317, 3, 2, 2, 2, 320, 323, 3, 2, 2, 2, 321, 319, 3, 2, 2,
	2, 321, 322, 3, 2, 2, 2, 322, 324, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 324,
	325, 7, 15, 2, 2, 325, 67, 3, 2, 2, 2, 326, 327, 7, 31, 2, 2, 327, 328,
	7, 51, 2, 2, 328, 329, 7, 16, 2, 2, 329, 330, 7, 51, 2, 2, 330, 331, 7,
	17, 2, 2, 331, 69, 3, 2, 2, 2, 332, 333, 7, 42, 2, 2, 333, 71, 3, 2, 2,
	2, 334, 335, 9, 4, 2, 2, 335, 73, 3, 2, 2, 2, 336, 337, 9, 5, 2, 2, 337,
	75, 3, 2, 2, 2, 338, 339, 9, 6, 2, 2, 339, 77, 3, 2, 2, 2, 340, 341, 9,
	7, 2, 2, 341, 79, 3, 2, 2, 2, 342, 343, 9, 8, 2, 2, 343, 81, 3, 2, 2, 2,
	344, 345, 7, 51, 2, 2, 345, 347, 7, 4, 2, 2, 346, 344, 3, 2, 2, 2, 347,
	350, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351,
	3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 351, 352, 7, 51, 2, 2, 352, 83, 3, 2,
	2, 2, 31, 86, 89, 92, 95, 107, 111, 130, 141, 149, 162, 186, 194, 208,
	223, 228, 231, 234, 237, 240, 250, 256, 262, 268, 283, 293, 295, 308, 321,
	348,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "', '", "','", "'*)'", "'d'", "'D'", "'h'", "'H'", "'m'", "'M'", "'s'",
	"'S'", "'{'", "'}'", "') = \"'", "'\"'", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	"", "", "", "", "", "'>'", "'>='", "'<'", "'<='", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "SELECTKEYWORD",
	"WHEREKEYWORD", "WINDOWKEYWORD", "GROUPBYKEYWORD", "HAVINGKEYWORD", "ALLWORD",
	"ASWORD", "SUMKEYWORD", "MINKEYWORD", "MAXKEYWORD", "AVGKEYWORD", "MEDIANWORD",
	"COUNTWORD", "TAGWORD", "FILEWORD", "INWORD", "PLUS", "MINUS", "TIMES",
	"DIVIDED", "LPAREN", "RPAREN", "AND", "OR", "NOT", "TRUE", "FALSE", "GT",
	"GE", "LT", "LE", "EQ", "NUMBER", "STRING", "DECITIMES", "IDENTIFIER",
	"WS",
}

var ruleNames = []string{
	"parse", "aggregateSelections", "selectAll", "selectElement", "selectDefault",
	"mathematicalSelection", "leftParen", "rightParen", "selectFunction", "selectSum",
	"selectMin", "selectMax", "selectAvg", "selectMedian", "selectCount", "countTag",
	"countNorTIMES", "groupByFunction", "whereFunction", "havingFunction",
	"windowFunction", "windowmode", "allMode", "timeMode", "valueMode", "tag",
	"days", "hours", "minutes", "seconds", "expression", "inexpressiontag",
	"inexpressionmetric", "hastag", "notnode", "endnode", "comparator", "binary",
	"mathematicalOperation", "boolToken", "list",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BitflowQueryParser struct {
	*antlr.BaseParser
}

func NewBitflowQueryParser(input antlr.TokenStream) *BitflowQueryParser {
	this := new(BitflowQueryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "BitflowQuery.g4"

	return this
}

// BitflowQueryParser tokens.
const (
	BitflowQueryParserEOF            = antlr.TokenEOF
	BitflowQueryParserT__0           = 1
	BitflowQueryParserT__1           = 2
	BitflowQueryParserT__2           = 3
	BitflowQueryParserT__3           = 4
	BitflowQueryParserT__4           = 5
	BitflowQueryParserT__5           = 6
	BitflowQueryParserT__6           = 7
	BitflowQueryParserT__7           = 8
	BitflowQueryParserT__8           = 9
	BitflowQueryParserT__9           = 10
	BitflowQueryParserT__10          = 11
	BitflowQueryParserT__11          = 12
	BitflowQueryParserT__12          = 13
	BitflowQueryParserT__13          = 14
	BitflowQueryParserT__14          = 15
	BitflowQueryParserSELECTKEYWORD  = 16
	BitflowQueryParserWHEREKEYWORD   = 17
	BitflowQueryParserWINDOWKEYWORD  = 18
	BitflowQueryParserGROUPBYKEYWORD = 19
	BitflowQueryParserHAVINGKEYWORD  = 20
	BitflowQueryParserALLWORD        = 21
	BitflowQueryParserASWORD         = 22
	BitflowQueryParserSUMKEYWORD     = 23
	BitflowQueryParserMINKEYWORD     = 24
	BitflowQueryParserMAXKEYWORD     = 25
	BitflowQueryParserAVGKEYWORD     = 26
	BitflowQueryParserMEDIANWORD     = 27
	BitflowQueryParserCOUNTWORD      = 28
	BitflowQueryParserTAGWORD        = 29
	BitflowQueryParserFILEWORD       = 30
	BitflowQueryParserINWORD         = 31
	BitflowQueryParserPLUS           = 32
	BitflowQueryParserMINUS          = 33
	BitflowQueryParserTIMES          = 34
	BitflowQueryParserDIVIDED        = 35
	BitflowQueryParserLPAREN         = 36
	BitflowQueryParserRPAREN         = 37
	BitflowQueryParserAND            = 38
	BitflowQueryParserOR             = 39
	BitflowQueryParserNOT            = 40
	BitflowQueryParserTRUE           = 41
	BitflowQueryParserFALSE          = 42
	BitflowQueryParserGT             = 43
	BitflowQueryParserGE             = 44
	BitflowQueryParserLT             = 45
	BitflowQueryParserLE             = 46
	BitflowQueryParserEQ             = 47
	BitflowQueryParserNUMBER         = 48
	BitflowQueryParserSTRING         = 49
	BitflowQueryParserDECITIMES      = 50
	BitflowQueryParserIDENTIFIER     = 51
	BitflowQueryParserWS             = 52
)

// BitflowQueryParser rules.
const (
	BitflowQueryParserRULE_parse                 = 0
	BitflowQueryParserRULE_aggregateSelections   = 1
	BitflowQueryParserRULE_selectAll             = 2
	BitflowQueryParserRULE_selectElement         = 3
	BitflowQueryParserRULE_selectDefault         = 4
	BitflowQueryParserRULE_mathematicalSelection = 5
	BitflowQueryParserRULE_leftParen             = 6
	BitflowQueryParserRULE_rightParen            = 7
	BitflowQueryParserRULE_selectFunction        = 8
	BitflowQueryParserRULE_selectSum             = 9
	BitflowQueryParserRULE_selectMin             = 10
	BitflowQueryParserRULE_selectMax             = 11
	BitflowQueryParserRULE_selectAvg             = 12
	BitflowQueryParserRULE_selectMedian          = 13
	BitflowQueryParserRULE_selectCount           = 14
	BitflowQueryParserRULE_countTag              = 15
	BitflowQueryParserRULE_countNorTIMES         = 16
	BitflowQueryParserRULE_groupByFunction       = 17
	BitflowQueryParserRULE_whereFunction         = 18
	BitflowQueryParserRULE_havingFunction        = 19
	BitflowQueryParserRULE_windowFunction        = 20
	BitflowQueryParserRULE_windowmode            = 21
	BitflowQueryParserRULE_allMode               = 22
	BitflowQueryParserRULE_timeMode              = 23
	BitflowQueryParserRULE_valueMode             = 24
	BitflowQueryParserRULE_tag                   = 25
	BitflowQueryParserRULE_days                  = 26
	BitflowQueryParserRULE_hours                 = 27
	BitflowQueryParserRULE_minutes               = 28
	BitflowQueryParserRULE_seconds               = 29
	BitflowQueryParserRULE_expression            = 30
	BitflowQueryParserRULE_inexpressiontag       = 31
	BitflowQueryParserRULE_inexpressionmetric    = 32
	BitflowQueryParserRULE_hastag                = 33
	BitflowQueryParserRULE_notnode               = 34
	BitflowQueryParserRULE_endnode               = 35
	BitflowQueryParserRULE_comparator            = 36
	BitflowQueryParserRULE_binary                = 37
	BitflowQueryParserRULE_mathematicalOperation = 38
	BitflowQueryParserRULE_boolToken             = 39
	BitflowQueryParserRULE_list                  = 40
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) AggregateSelections() IAggregateSelectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregateSelectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggregateSelectionsContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserEOF, 0)
}

func (s *ParseContext) WhereFunction() IWhereFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereFunctionContext)
}

func (s *ParseContext) GroupByFunction() IGroupByFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupByFunctionContext)
}

func (s *ParseContext) WindowFunction() IWindowFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWindowFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWindowFunctionContext)
}

func (s *ParseContext) HavingFunction() IHavingFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHavingFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHavingFunctionContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BitflowQueryParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.AggregateSelections()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserWHEREKEYWORD {
		{
			p.SetState(83)
			p.WhereFunction()
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserGROUPBYKEYWORD {
		{
			p.SetState(86)
			p.GroupByFunction()
		}

	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserWINDOWKEYWORD {
		{
			p.SetState(89)
			p.WindowFunction()
		}

	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserHAVINGKEYWORD {
		{
			p.SetState(92)
			p.HavingFunction()
		}

	}
	{
		p.SetState(95)
		p.Match(BitflowQueryParserEOF)
	}

	return localctx
}

// IAggregateSelectionsContext is an interface to support dynamic dispatch.
type IAggregateSelectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregateSelectionsContext differentiates from other interfaces.
	IsAggregateSelectionsContext()
}

type AggregateSelectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateSelectionsContext() *AggregateSelectionsContext {
	var p = new(AggregateSelectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_aggregateSelections
	return p
}

func (*AggregateSelectionsContext) IsAggregateSelectionsContext() {}

func NewAggregateSelectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateSelectionsContext {
	var p = new(AggregateSelectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_aggregateSelections

	return p
}

func (s *AggregateSelectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateSelectionsContext) SELECTKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSELECTKEYWORD, 0)
}

func (s *AggregateSelectionsContext) SelectAll() ISelectAllContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectAllContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectAllContext)
}

func (s *AggregateSelectionsContext) AllSelectElement() []ISelectElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectElementContext)(nil)).Elem())
	var tst = make([]ISelectElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectElementContext)
		}
	}

	return tst
}

func (s *AggregateSelectionsContext) SelectElement(i int) ISelectElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *AggregateSelectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateSelectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateSelectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterAggregateSelections(s)
	}
}

func (s *AggregateSelectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitAggregateSelections(s)
	}
}

func (s *AggregateSelectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitAggregateSelections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) AggregateSelections() (localctx IAggregateSelectionsContext) {
	localctx = NewAggregateSelectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BitflowQueryParserRULE_aggregateSelections)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(BitflowQueryParserSELECTKEYWORD)
		}
		{
			p.SetState(98)
			p.SelectAll()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(BitflowQueryParserSELECTKEYWORD)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(100)
					p.SelectElement()
				}
				{
					p.SetState(101)
					p.Match(BitflowQueryParserT__0)
				}

			}
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}
		{
			p.SetState(108)
			p.SelectElement()
		}

	}

	return localctx
}

// ISelectAllContext is an interface to support dynamic dispatch.
type ISelectAllContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectAllContext differentiates from other interfaces.
	IsSelectAllContext()
}

type SelectAllContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectAllContext() *SelectAllContext {
	var p = new(SelectAllContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectAll
	return p
}

func (*SelectAllContext) IsSelectAllContext() {}

func NewSelectAllContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectAllContext {
	var p = new(SelectAllContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectAll

	return p
}

func (s *SelectAllContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectAllContext) ALLWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserALLWORD, 0)
}

func (s *SelectAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectAllContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectAll(s)
	}
}

func (s *SelectAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectAll(s)
	}
}

func (s *SelectAllContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectAll(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectAll() (localctx ISelectAllContext) {
	localctx = NewSelectAllContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BitflowQueryParserRULE_selectAll)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowQueryParserALLWORD || _la == BitflowQueryParserTIMES) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) MathematicalSelection() IMathematicalSelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathematicalSelectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathematicalSelectionContext)
}

func (s *SelectElementContext) ASWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserASWORD, 0)
}

func (s *SelectElementContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectElementContext) SelectFunction() ISelectFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectFunctionContext)
}

func (s *SelectElementContext) SelectDefault() ISelectDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectDefaultContext)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectElement(s)
	}
}

func (s *SelectElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectElement(s)
	}
}

func (s *SelectElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectElement() (localctx ISelectElementContext) {
	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BitflowQueryParserRULE_selectElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.mathematicalSelection(0)
		}
		{
			p.SetState(114)
			p.Match(BitflowQueryParserASWORD)
		}
		{
			p.SetState(115)
			p.Match(BitflowQueryParserSTRING)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.mathematicalSelection(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.SelectFunction()
		}
		{
			p.SetState(119)
			p.Match(BitflowQueryParserASWORD)
		}
		{
			p.SetState(120)
			p.Match(BitflowQueryParserSTRING)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.SelectFunction()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.SelectDefault()
		}
		{
			p.SetState(124)
			p.Match(BitflowQueryParserASWORD)
		}
		{
			p.SetState(125)
			p.Match(BitflowQueryParserSTRING)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.SelectDefault()
		}

	}

	return localctx
}

// ISelectDefaultContext is an interface to support dynamic dispatch.
type ISelectDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectDefaultContext differentiates from other interfaces.
	IsSelectDefaultContext()
}

type SelectDefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectDefaultContext() *SelectDefaultContext {
	var p = new(SelectDefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectDefault
	return p
}

func (*SelectDefaultContext) IsSelectDefaultContext() {}

func NewSelectDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectDefaultContext {
	var p = new(SelectDefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectDefault

	return p
}

func (s *SelectDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectDefaultContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectDefaultContext) DECITIMES() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserDECITIMES, 0)
}

func (s *SelectDefaultContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *SelectDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectDefault(s)
	}
}

func (s *SelectDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectDefault(s)
	}
}

func (s *SelectDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectDefault() (localctx ISelectDefaultContext) {
	localctx = NewSelectDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BitflowQueryParserRULE_selectDefault)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(BitflowQueryParserNUMBER-48))|(1<<(BitflowQueryParserSTRING-48))|(1<<(BitflowQueryParserDECITIMES-48)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathematicalSelectionContext is an interface to support dynamic dispatch.
type IMathematicalSelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathematicalSelectionContext differentiates from other interfaces.
	IsMathematicalSelectionContext()
}

type MathematicalSelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathematicalSelectionContext() *MathematicalSelectionContext {
	var p = new(MathematicalSelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_mathematicalSelection
	return p
}

func (*MathematicalSelectionContext) IsMathematicalSelectionContext() {}

func NewMathematicalSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathematicalSelectionContext {
	var p = new(MathematicalSelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_mathematicalSelection

	return p
}

func (s *MathematicalSelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathematicalSelectionContext) LeftParen() ILeftParenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeftParenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeftParenContext)
}

func (s *MathematicalSelectionContext) AllMathematicalSelection() []IMathematicalSelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathematicalSelectionContext)(nil)).Elem())
	var tst = make([]IMathematicalSelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathematicalSelectionContext)
		}
	}

	return tst
}

func (s *MathematicalSelectionContext) MathematicalSelection(i int) IMathematicalSelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathematicalSelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathematicalSelectionContext)
}

func (s *MathematicalSelectionContext) RightParen() IRightParenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRightParenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRightParenContext)
}

func (s *MathematicalSelectionContext) SelectFunction() ISelectFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectFunctionContext)
}

func (s *MathematicalSelectionContext) SelectDefault() ISelectDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectDefaultContext)
}

func (s *MathematicalSelectionContext) MathematicalOperation() IMathematicalOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathematicalOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathematicalOperationContext)
}

func (s *MathematicalSelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathematicalSelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathematicalSelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterMathematicalSelection(s)
	}
}

func (s *MathematicalSelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitMathematicalSelection(s)
	}
}

func (s *MathematicalSelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitMathematicalSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) MathematicalSelection() (localctx IMathematicalSelectionContext) {
	return p.mathematicalSelection(0)
}

func (p *BitflowQueryParser) mathematicalSelection(_p int) (localctx IMathematicalSelectionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathematicalSelectionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathematicalSelectionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, BitflowQueryParserRULE_mathematicalSelection, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BitflowQueryParserLPAREN:
		{
			p.SetState(133)
			p.LeftParen()
		}
		{
			p.SetState(134)
			p.mathematicalSelection(0)
		}
		{
			p.SetState(135)
			p.RightParen()
		}

	case BitflowQueryParserSUMKEYWORD, BitflowQueryParserMINKEYWORD, BitflowQueryParserMAXKEYWORD, BitflowQueryParserAVGKEYWORD, BitflowQueryParserMEDIANWORD, BitflowQueryParserCOUNTWORD:
		{
			p.SetState(137)
			p.SelectFunction()
		}

	case BitflowQueryParserNUMBER, BitflowQueryParserSTRING, BitflowQueryParserDECITIMES:
		{
			p.SetState(138)
			p.SelectDefault()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMathematicalSelectionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, BitflowQueryParserRULE_mathematicalSelection)
			p.SetState(141)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(142)
				p.MathematicalOperation()
			}
			{
				p.SetState(143)
				p.mathematicalSelection(4)
			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// ILeftParenContext is an interface to support dynamic dispatch.
type ILeftParenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeftParenContext differentiates from other interfaces.
	IsLeftParenContext()
}

type LeftParenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftParenContext() *LeftParenContext {
	var p = new(LeftParenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_leftParen
	return p
}

func (*LeftParenContext) IsLeftParenContext() {}

func NewLeftParenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftParenContext {
	var p = new(LeftParenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_leftParen

	return p
}

func (s *LeftParenContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserLPAREN, 0)
}

func (s *LeftParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftParenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterLeftParen(s)
	}
}

func (s *LeftParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitLeftParen(s)
	}
}

func (s *LeftParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitLeftParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) LeftParen() (localctx ILeftParenContext) {
	localctx = NewLeftParenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BitflowQueryParserRULE_leftParen)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(BitflowQueryParserLPAREN)
	}

	return localctx
}

// IRightParenContext is an interface to support dynamic dispatch.
type IRightParenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRightParenContext differentiates from other interfaces.
	IsRightParenContext()
}

type RightParenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightParenContext() *RightParenContext {
	var p = new(RightParenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_rightParen
	return p
}

func (*RightParenContext) IsRightParenContext() {}

func NewRightParenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightParenContext {
	var p = new(RightParenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_rightParen

	return p
}

func (s *RightParenContext) GetParser() antlr.Parser { return s.parser }

func (s *RightParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserRPAREN, 0)
}

func (s *RightParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightParenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterRightParen(s)
	}
}

func (s *RightParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitRightParen(s)
	}
}

func (s *RightParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitRightParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) RightParen() (localctx IRightParenContext) {
	localctx = NewRightParenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BitflowQueryParserRULE_rightParen)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectFunctionContext is an interface to support dynamic dispatch.
type ISelectFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectFunctionContext differentiates from other interfaces.
	IsSelectFunctionContext()
}

type SelectFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFunctionContext() *SelectFunctionContext {
	var p = new(SelectFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectFunction
	return p
}

func (*SelectFunctionContext) IsSelectFunctionContext() {}

func NewSelectFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFunctionContext {
	var p = new(SelectFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectFunction

	return p
}

func (s *SelectFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFunctionContext) SelectSum() ISelectSumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectSumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectSumContext)
}

func (s *SelectFunctionContext) SelectMin() ISelectMinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectMinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectMinContext)
}

func (s *SelectFunctionContext) SelectMax() ISelectMaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectMaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectMaxContext)
}

func (s *SelectFunctionContext) SelectAvg() ISelectAvgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectAvgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectAvgContext)
}

func (s *SelectFunctionContext) SelectMedian() ISelectMedianContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectMedianContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectMedianContext)
}

func (s *SelectFunctionContext) SelectCount() ISelectCountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectCountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectCountContext)
}

func (s *SelectFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectFunction(s)
	}
}

func (s *SelectFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectFunction(s)
	}
}

func (s *SelectFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectFunction() (localctx ISelectFunctionContext) {
	localctx = NewSelectFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BitflowQueryParserRULE_selectFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BitflowQueryParserSUMKEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.SelectSum()
		}

	case BitflowQueryParserMINKEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.SelectMin()
		}

	case BitflowQueryParserMAXKEYWORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.SelectMax()
		}

	case BitflowQueryParserAVGKEYWORD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(157)
			p.SelectAvg()
		}

	case BitflowQueryParserMEDIANWORD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)
			p.SelectMedian()
		}

	case BitflowQueryParserCOUNTWORD:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(159)
			p.SelectCount()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectSumContext is an interface to support dynamic dispatch.
type ISelectSumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectSumContext differentiates from other interfaces.
	IsSelectSumContext()
}

type SelectSumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectSumContext() *SelectSumContext {
	var p = new(SelectSumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectSum
	return p
}

func (*SelectSumContext) IsSelectSumContext() {}

func NewSelectSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectSumContext {
	var p = new(SelectSumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectSum

	return p
}

func (s *SelectSumContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectSumContext) SUMKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSUMKEYWORD, 0)
}

func (s *SelectSumContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectSumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectSum(s)
	}
}

func (s *SelectSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectSum(s)
	}
}

func (s *SelectSumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectSum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectSum() (localctx ISelectSumContext) {
	localctx = NewSelectSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BitflowQueryParserRULE_selectSum)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(BitflowQueryParserSUMKEYWORD)
	}
	{
		p.SetState(163)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(164)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectMinContext is an interface to support dynamic dispatch.
type ISelectMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectMinContext differentiates from other interfaces.
	IsSelectMinContext()
}

type SelectMinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectMinContext() *SelectMinContext {
	var p = new(SelectMinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectMin
	return p
}

func (*SelectMinContext) IsSelectMinContext() {}

func NewSelectMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectMinContext {
	var p = new(SelectMinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectMin

	return p
}

func (s *SelectMinContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectMinContext) MINKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserMINKEYWORD, 0)
}

func (s *SelectMinContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectMinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectMin(s)
	}
}

func (s *SelectMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectMin(s)
	}
}

func (s *SelectMinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectMin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectMin() (localctx ISelectMinContext) {
	localctx = NewSelectMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BitflowQueryParserRULE_selectMin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(BitflowQueryParserMINKEYWORD)
	}
	{
		p.SetState(167)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(168)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectMaxContext is an interface to support dynamic dispatch.
type ISelectMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectMaxContext differentiates from other interfaces.
	IsSelectMaxContext()
}

type SelectMaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectMaxContext() *SelectMaxContext {
	var p = new(SelectMaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectMax
	return p
}

func (*SelectMaxContext) IsSelectMaxContext() {}

func NewSelectMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectMaxContext {
	var p = new(SelectMaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectMax

	return p
}

func (s *SelectMaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectMaxContext) MAXKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserMAXKEYWORD, 0)
}

func (s *SelectMaxContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectMaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectMax(s)
	}
}

func (s *SelectMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectMax(s)
	}
}

func (s *SelectMaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectMax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectMax() (localctx ISelectMaxContext) {
	localctx = NewSelectMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BitflowQueryParserRULE_selectMax)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(BitflowQueryParserMAXKEYWORD)
	}
	{
		p.SetState(171)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(172)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectAvgContext is an interface to support dynamic dispatch.
type ISelectAvgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectAvgContext differentiates from other interfaces.
	IsSelectAvgContext()
}

type SelectAvgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectAvgContext() *SelectAvgContext {
	var p = new(SelectAvgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectAvg
	return p
}

func (*SelectAvgContext) IsSelectAvgContext() {}

func NewSelectAvgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectAvgContext {
	var p = new(SelectAvgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectAvg

	return p
}

func (s *SelectAvgContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectAvgContext) AVGKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserAVGKEYWORD, 0)
}

func (s *SelectAvgContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectAvgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectAvgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectAvgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectAvg(s)
	}
}

func (s *SelectAvgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectAvg(s)
	}
}

func (s *SelectAvgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectAvg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectAvg() (localctx ISelectAvgContext) {
	localctx = NewSelectAvgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BitflowQueryParserRULE_selectAvg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(BitflowQueryParserAVGKEYWORD)
	}
	{
		p.SetState(175)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(176)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectMedianContext is an interface to support dynamic dispatch.
type ISelectMedianContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectMedianContext differentiates from other interfaces.
	IsSelectMedianContext()
}

type SelectMedianContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectMedianContext() *SelectMedianContext {
	var p = new(SelectMedianContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectMedian
	return p
}

func (*SelectMedianContext) IsSelectMedianContext() {}

func NewSelectMedianContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectMedianContext {
	var p = new(SelectMedianContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectMedian

	return p
}

func (s *SelectMedianContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectMedianContext) MEDIANWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserMEDIANWORD, 0)
}

func (s *SelectMedianContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *SelectMedianContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectMedianContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectMedianContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectMedian(s)
	}
}

func (s *SelectMedianContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectMedian(s)
	}
}

func (s *SelectMedianContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectMedian(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectMedian() (localctx ISelectMedianContext) {
	localctx = NewSelectMedianContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BitflowQueryParserRULE_selectMedian)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(BitflowQueryParserMEDIANWORD)
	}
	{
		p.SetState(179)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(180)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ISelectCountContext is an interface to support dynamic dispatch.
type ISelectCountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectCountContext differentiates from other interfaces.
	IsSelectCountContext()
}

type SelectCountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectCountContext() *SelectCountContext {
	var p = new(SelectCountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_selectCount
	return p
}

func (*SelectCountContext) IsSelectCountContext() {}

func NewSelectCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectCountContext {
	var p = new(SelectCountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_selectCount

	return p
}

func (s *SelectCountContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectCountContext) CountTag() ICountTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountTagContext)
}

func (s *SelectCountContext) CountNorTIMES() ICountNorTIMESContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountNorTIMESContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountNorTIMESContext)
}

func (s *SelectCountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectCountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectCountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSelectCount(s)
	}
}

func (s *SelectCountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSelectCount(s)
	}
}

func (s *SelectCountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSelectCount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) SelectCount() (localctx ISelectCountContext) {
	localctx = NewSelectCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BitflowQueryParserRULE_selectCount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.CountTag()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.CountNorTIMES()
		}

	}

	return localctx
}

// ICountTagContext is an interface to support dynamic dispatch.
type ICountTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountTagContext differentiates from other interfaces.
	IsCountTagContext()
}

type CountTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountTagContext() *CountTagContext {
	var p = new(CountTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_countTag
	return p
}

func (*CountTagContext) IsCountTagContext() {}

func NewCountTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountTagContext {
	var p = new(CountTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_countTag

	return p
}

func (s *CountTagContext) GetParser() antlr.Parser { return s.parser }

func (s *CountTagContext) COUNTWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserCOUNTWORD, 0)
}

func (s *CountTagContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BitflowQueryParserSTRING)
}

func (s *CountTagContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, i)
}

func (s *CountTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterCountTag(s)
	}
}

func (s *CountTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitCountTag(s)
	}
}

func (s *CountTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitCountTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) CountTag() (localctx ICountTagContext) {
	localctx = NewCountTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BitflowQueryParserRULE_countTag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(BitflowQueryParserCOUNTWORD)
	}
	{
		p.SetState(187)
		p.Match(BitflowQueryParserSTRING)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowQueryParserT__1 {
		{
			p.SetState(188)
			p.Match(BitflowQueryParserT__1)
		}
		{
			p.SetState(189)
			p.Match(BitflowQueryParserSTRING)
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
		p.Match(BitflowQueryParserRPAREN)
	}

	return localctx
}

// ICountNorTIMESContext is an interface to support dynamic dispatch.
type ICountNorTIMESContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountNorTIMESContext differentiates from other interfaces.
	IsCountNorTIMESContext()
}

type CountNorTIMESContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountNorTIMESContext() *CountNorTIMESContext {
	var p = new(CountNorTIMESContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_countNorTIMES
	return p
}

func (*CountNorTIMESContext) IsCountNorTIMESContext() {}

func NewCountNorTIMESContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountNorTIMESContext {
	var p = new(CountNorTIMESContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_countNorTIMES

	return p
}

func (s *CountNorTIMESContext) GetParser() antlr.Parser { return s.parser }

func (s *CountNorTIMESContext) COUNTWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserCOUNTWORD, 0)
}

func (s *CountNorTIMESContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountNorTIMESContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountNorTIMESContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterCountNorTIMES(s)
	}
}

func (s *CountNorTIMESContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitCountNorTIMES(s)
	}
}

func (s *CountNorTIMESContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitCountNorTIMES(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) CountNorTIMES() (localctx ICountNorTIMESContext) {
	localctx = NewCountNorTIMESContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BitflowQueryParserRULE_countNorTIMES)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(BitflowQueryParserCOUNTWORD)
	}
	{
		p.SetState(198)
		p.Match(BitflowQueryParserT__2)
	}

	return localctx
}

// IGroupByFunctionContext is an interface to support dynamic dispatch.
type IGroupByFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByFunctionContext differentiates from other interfaces.
	IsGroupByFunctionContext()
}

type GroupByFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByFunctionContext() *GroupByFunctionContext {
	var p = new(GroupByFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_groupByFunction
	return p
}

func (*GroupByFunctionContext) IsGroupByFunctionContext() {}

func NewGroupByFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByFunctionContext {
	var p = new(GroupByFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_groupByFunction

	return p
}

func (s *GroupByFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByFunctionContext) GROUPBYKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserGROUPBYKEYWORD, 0)
}

func (s *GroupByFunctionContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *GroupByFunctionContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *GroupByFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterGroupByFunction(s)
	}
}

func (s *GroupByFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitGroupByFunction(s)
	}
}

func (s *GroupByFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitGroupByFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) GroupByFunction() (localctx IGroupByFunctionContext) {
	localctx = NewGroupByFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BitflowQueryParserRULE_groupByFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(BitflowQueryParserGROUPBYKEYWORD)
	}
	{
		p.SetState(201)
		p.Tag()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowQueryParserT__0 {
		{
			p.SetState(202)
			p.Match(BitflowQueryParserT__0)
		}
		{
			p.SetState(203)
			p.Tag()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWhereFunctionContext is an interface to support dynamic dispatch.
type IWhereFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereFunctionContext differentiates from other interfaces.
	IsWhereFunctionContext()
}

type WhereFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereFunctionContext() *WhereFunctionContext {
	var p = new(WhereFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_whereFunction
	return p
}

func (*WhereFunctionContext) IsWhereFunctionContext() {}

func NewWhereFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereFunctionContext {
	var p = new(WhereFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_whereFunction

	return p
}

func (s *WhereFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereFunctionContext) WHEREKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserWHEREKEYWORD, 0)
}

func (s *WhereFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhereFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterWhereFunction(s)
	}
}

func (s *WhereFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitWhereFunction(s)
	}
}

func (s *WhereFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitWhereFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) WhereFunction() (localctx IWhereFunctionContext) {
	localctx = NewWhereFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BitflowQueryParserRULE_whereFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(BitflowQueryParserWHEREKEYWORD)
	}
	{
		p.SetState(210)
		p.expression(0)
	}

	return localctx
}

// IHavingFunctionContext is an interface to support dynamic dispatch.
type IHavingFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHavingFunctionContext differentiates from other interfaces.
	IsHavingFunctionContext()
}

type HavingFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingFunctionContext() *HavingFunctionContext {
	var p = new(HavingFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_havingFunction
	return p
}

func (*HavingFunctionContext) IsHavingFunctionContext() {}

func NewHavingFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingFunctionContext {
	var p = new(HavingFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_havingFunction

	return p
}

func (s *HavingFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingFunctionContext) HAVINGKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserHAVINGKEYWORD, 0)
}

func (s *HavingFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HavingFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterHavingFunction(s)
	}
}

func (s *HavingFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitHavingFunction(s)
	}
}

func (s *HavingFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitHavingFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) HavingFunction() (localctx IHavingFunctionContext) {
	localctx = NewHavingFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BitflowQueryParserRULE_havingFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(BitflowQueryParserHAVINGKEYWORD)
	}
	{
		p.SetState(213)
		p.expression(0)
	}

	return localctx
}

// IWindowFunctionContext is an interface to support dynamic dispatch.
type IWindowFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowFunctionContext differentiates from other interfaces.
	IsWindowFunctionContext()
}

type WindowFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowFunctionContext() *WindowFunctionContext {
	var p = new(WindowFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_windowFunction
	return p
}

func (*WindowFunctionContext) IsWindowFunctionContext() {}

func NewWindowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowFunctionContext {
	var p = new(WindowFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_windowFunction

	return p
}

func (s *WindowFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowFunctionContext) WINDOWKEYWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserWINDOWKEYWORD, 0)
}

func (s *WindowFunctionContext) Windowmode() IWindowmodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWindowmodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWindowmodeContext)
}

func (s *WindowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterWindowFunction(s)
	}
}

func (s *WindowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitWindowFunction(s)
	}
}

func (s *WindowFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitWindowFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) WindowFunction() (localctx IWindowFunctionContext) {
	localctx = NewWindowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BitflowQueryParserRULE_windowFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(BitflowQueryParserWINDOWKEYWORD)
	}
	{
		p.SetState(216)
		p.Windowmode()
	}

	return localctx
}

// IWindowmodeContext is an interface to support dynamic dispatch.
type IWindowmodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowmodeContext differentiates from other interfaces.
	IsWindowmodeContext()
}

type WindowmodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowmodeContext() *WindowmodeContext {
	var p = new(WindowmodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_windowmode
	return p
}

func (*WindowmodeContext) IsWindowmodeContext() {}

func NewWindowmodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowmodeContext {
	var p = new(WindowmodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_windowmode

	return p
}

func (s *WindowmodeContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowmodeContext) AllMode() IAllModeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllModeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllModeContext)
}

func (s *WindowmodeContext) ValueMode() IValueModeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueModeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueModeContext)
}

func (s *WindowmodeContext) TimeMode() ITimeModeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeModeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeModeContext)
}

func (s *WindowmodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowmodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowmodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterWindowmode(s)
	}
}

func (s *WindowmodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitWindowmode(s)
	}
}

func (s *WindowmodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitWindowmode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Windowmode() (localctx IWindowmodeContext) {
	localctx = NewWindowmodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BitflowQueryParserRULE_windowmode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.AllMode()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.ValueMode()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.TimeMode()
		}

	}

	return localctx
}

// IAllModeContext is an interface to support dynamic dispatch.
type IAllModeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllModeContext differentiates from other interfaces.
	IsAllModeContext()
}

type AllModeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllModeContext() *AllModeContext {
	var p = new(AllModeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_allMode
	return p
}

func (*AllModeContext) IsAllModeContext() {}

func NewAllModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllModeContext {
	var p = new(AllModeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_allMode

	return p
}

func (s *AllModeContext) GetParser() antlr.Parser { return s.parser }

func (s *AllModeContext) ALLWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserALLWORD, 0)
}

func (s *AllModeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllModeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllModeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterAllMode(s)
	}
}

func (s *AllModeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitAllMode(s)
	}
}

func (s *AllModeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitAllMode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) AllMode() (localctx IAllModeContext) {
	localctx = NewAllModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BitflowQueryParserRULE_allMode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(BitflowQueryParserALLWORD)
	}

	return localctx
}

// ITimeModeContext is an interface to support dynamic dispatch.
type ITimeModeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeModeContext differentiates from other interfaces.
	IsTimeModeContext()
}

type TimeModeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeModeContext() *TimeModeContext {
	var p = new(TimeModeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_timeMode
	return p
}

func (*TimeModeContext) IsTimeModeContext() {}

func NewTimeModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeModeContext {
	var p = new(TimeModeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_timeMode

	return p
}

func (s *TimeModeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeModeContext) FILEWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserFILEWORD, 0)
}

func (s *TimeModeContext) Days() IDaysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDaysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDaysContext)
}

func (s *TimeModeContext) Hours() IHoursContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHoursContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHoursContext)
}

func (s *TimeModeContext) Minutes() IMinutesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinutesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinutesContext)
}

func (s *TimeModeContext) Seconds() ISecondsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondsContext)
}

func (s *TimeModeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeModeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeModeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterTimeMode(s)
	}
}

func (s *TimeModeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitTimeMode(s)
	}
}

func (s *TimeModeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitTimeMode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) TimeMode() (localctx ITimeModeContext) {
	localctx = NewTimeModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BitflowQueryParserRULE_timeMode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserFILEWORD {
		{
			p.SetState(225)
			p.Match(BitflowQueryParserFILEWORD)
		}

	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(228)
			p.Days()
		}

	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(231)
			p.Hours()
		}

	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(234)
			p.Minutes()
		}

	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BitflowQueryParserNUMBER {
		{
			p.SetState(237)
			p.Seconds()
		}

	}

	return localctx
}

// IValueModeContext is an interface to support dynamic dispatch.
type IValueModeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueModeContext differentiates from other interfaces.
	IsValueModeContext()
}

type ValueModeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueModeContext() *ValueModeContext {
	var p = new(ValueModeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_valueMode
	return p
}

func (*ValueModeContext) IsValueModeContext() {}

func NewValueModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueModeContext {
	var p = new(ValueModeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_valueMode

	return p
}

func (s *ValueModeContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueModeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *ValueModeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueModeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueModeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterValueMode(s)
	}
}

func (s *ValueModeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitValueMode(s)
	}
}

func (s *ValueModeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitValueMode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) ValueMode() (localctx IValueModeContext) {
	localctx = NewValueModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BitflowQueryParserRULE_valueMode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(BitflowQueryParserNUMBER)
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitTag(s)
	}
}

func (s *TagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BitflowQueryParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(BitflowQueryParserSTRING)
	}

	return localctx
}

// IDaysContext is an interface to support dynamic dispatch.
type IDaysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDaysContext differentiates from other interfaces.
	IsDaysContext()
}

type DaysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDaysContext() *DaysContext {
	var p = new(DaysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_days
	return p
}

func (*DaysContext) IsDaysContext() {}

func NewDaysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DaysContext {
	var p = new(DaysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_days

	return p
}

func (s *DaysContext) GetParser() antlr.Parser { return s.parser }

func (s *DaysContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *DaysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DaysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DaysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterDays(s)
	}
}

func (s *DaysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitDays(s)
	}
}

func (s *DaysContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitDays(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Days() (localctx IDaysContext) {
	localctx = NewDaysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BitflowQueryParserRULE_days)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(245)
			p.Match(BitflowQueryParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(247)
			p.Match(BitflowQueryParserT__4)
		}

	}

	return localctx
}

// IHoursContext is an interface to support dynamic dispatch.
type IHoursContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHoursContext differentiates from other interfaces.
	IsHoursContext()
}

type HoursContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHoursContext() *HoursContext {
	var p = new(HoursContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_hours
	return p
}

func (*HoursContext) IsHoursContext() {}

func NewHoursContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HoursContext {
	var p = new(HoursContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_hours

	return p
}

func (s *HoursContext) GetParser() antlr.Parser { return s.parser }

func (s *HoursContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *HoursContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HoursContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HoursContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterHours(s)
	}
}

func (s *HoursContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitHours(s)
	}
}

func (s *HoursContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitHours(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Hours() (localctx IHoursContext) {
	localctx = NewHoursContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BitflowQueryParserRULE_hours)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(251)
			p.Match(BitflowQueryParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(253)
			p.Match(BitflowQueryParserT__6)
		}

	}

	return localctx
}

// IMinutesContext is an interface to support dynamic dispatch.
type IMinutesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinutesContext differentiates from other interfaces.
	IsMinutesContext()
}

type MinutesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinutesContext() *MinutesContext {
	var p = new(MinutesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_minutes
	return p
}

func (*MinutesContext) IsMinutesContext() {}

func NewMinutesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinutesContext {
	var p = new(MinutesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_minutes

	return p
}

func (s *MinutesContext) GetParser() antlr.Parser { return s.parser }

func (s *MinutesContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *MinutesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinutesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinutesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterMinutes(s)
	}
}

func (s *MinutesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitMinutes(s)
	}
}

func (s *MinutesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitMinutes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Minutes() (localctx IMinutesContext) {
	localctx = NewMinutesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BitflowQueryParserRULE_minutes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(257)
			p.Match(BitflowQueryParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(259)
			p.Match(BitflowQueryParserT__8)
		}

	}

	return localctx
}

// ISecondsContext is an interface to support dynamic dispatch.
type ISecondsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondsContext differentiates from other interfaces.
	IsSecondsContext()
}

type SecondsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondsContext() *SecondsContext {
	var p = new(SecondsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_seconds
	return p
}

func (*SecondsContext) IsSecondsContext() {}

func NewSecondsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondsContext {
	var p = new(SecondsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_seconds

	return p
}

func (s *SecondsContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondsContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *SecondsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterSeconds(s)
	}
}

func (s *SecondsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitSeconds(s)
	}
}

func (s *SecondsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitSeconds(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Seconds() (localctx ISecondsContext) {
	localctx = NewSecondsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BitflowQueryParserRULE_seconds)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(263)
			p.Match(BitflowQueryParserT__9)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Match(BitflowQueryParserNUMBER)
		}
		{
			p.SetState(265)
			p.Match(BitflowQueryParserT__10)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpressionContext
	right  IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserLPAREN, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserRPAREN, 0)
}

func (s *ExpressionContext) Notnode() INotnodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotnodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotnodeContext)
}

func (s *ExpressionContext) BoolToken() IBoolTokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolTokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolTokenContext)
}

func (s *ExpressionContext) Inexpressionmetric() IInexpressionmetricContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInexpressionmetricContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInexpressionmetricContext)
}

func (s *ExpressionContext) Inexpressiontag() IInexpressiontagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInexpressiontagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInexpressiontagContext)
}

func (s *ExpressionContext) Hastag() IHastagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHastagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHastagContext)
}

func (s *ExpressionContext) Endnode() IEndnodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndnodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndnodeContext)
}

func (s *ExpressionContext) Comparator() IComparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ExpressionContext) Binary() IBinaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *BitflowQueryParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, BitflowQueryParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(269)
			p.Match(BitflowQueryParserLPAREN)
		}
		{
			p.SetState(270)
			p.expression(0)
		}
		{
			p.SetState(271)
			p.Match(BitflowQueryParserRPAREN)
		}

	case 2:
		{
			p.SetState(273)
			p.Notnode()
		}
		{
			p.SetState(274)
			p.expression(8)
		}

	case 3:
		{
			p.SetState(276)
			p.BoolToken()
		}

	case 4:
		{
			p.SetState(277)
			p.Inexpressionmetric()
		}

	case 5:
		{
			p.SetState(278)
			p.Inexpressiontag()
		}

	case 6:
		{
			p.SetState(279)
			p.Hastag()
		}

	case 7:
		{
			p.SetState(280)
			p.Endnode()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, BitflowQueryParserRULE_expression)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(284)
					p.Comparator()
				}
				{
					p.SetState(285)

					var _x = p.expression(8)

					localctx.(*ExpressionContext).right = _x
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				localctx.(*ExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, BitflowQueryParserRULE_expression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(288)
					p.Binary()
				}
				{
					p.SetState(289)

					var _x = p.expression(7)

					localctx.(*ExpressionContext).right = _x
				}

			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IInexpressiontagContext is an interface to support dynamic dispatch.
type IInexpressiontagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInexpressiontagContext differentiates from other interfaces.
	IsInexpressiontagContext()
}

type InexpressiontagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInexpressiontagContext() *InexpressiontagContext {
	var p = new(InexpressiontagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_inexpressiontag
	return p
}

func (*InexpressiontagContext) IsInexpressiontagContext() {}

func NewInexpressiontagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InexpressiontagContext {
	var p = new(InexpressiontagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_inexpressiontag

	return p
}

func (s *InexpressiontagContext) GetParser() antlr.Parser { return s.parser }

func (s *InexpressiontagContext) TAGWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserTAGWORD, 0)
}

func (s *InexpressiontagContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BitflowQueryParserSTRING)
}

func (s *InexpressiontagContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, i)
}

func (s *InexpressiontagContext) INWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserINWORD, 0)
}

func (s *InexpressiontagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InexpressiontagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InexpressiontagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterInexpressiontag(s)
	}
}

func (s *InexpressiontagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitInexpressiontag(s)
	}
}

func (s *InexpressiontagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitInexpressiontag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Inexpressiontag() (localctx IInexpressiontagContext) {
	localctx = NewInexpressiontagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BitflowQueryParserRULE_inexpressiontag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(BitflowQueryParserTAGWORD)
	}
	{
		p.SetState(297)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(298)
		p.Match(BitflowQueryParserRPAREN)
	}
	{
		p.SetState(299)
		p.Match(BitflowQueryParserINWORD)
	}
	{
		p.SetState(300)
		p.Match(BitflowQueryParserT__11)
	}
	{
		p.SetState(301)
		p.Match(BitflowQueryParserSTRING)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowQueryParserT__1 {
		{
			p.SetState(302)
			p.Match(BitflowQueryParserT__1)
		}
		{
			p.SetState(303)
			p.Match(BitflowQueryParserSTRING)
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
		p.Match(BitflowQueryParserT__12)
	}

	return localctx
}

// IInexpressionmetricContext is an interface to support dynamic dispatch.
type IInexpressionmetricContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInexpressionmetricContext differentiates from other interfaces.
	IsInexpressionmetricContext()
}

type InexpressionmetricContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInexpressionmetricContext() *InexpressionmetricContext {
	var p = new(InexpressionmetricContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_inexpressionmetric
	return p
}

func (*InexpressionmetricContext) IsInexpressionmetricContext() {}

func NewInexpressionmetricContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InexpressionmetricContext {
	var p = new(InexpressionmetricContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_inexpressionmetric

	return p
}

func (s *InexpressionmetricContext) GetParser() antlr.Parser { return s.parser }

func (s *InexpressionmetricContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *InexpressionmetricContext) INWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserINWORD, 0)
}

func (s *InexpressionmetricContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(BitflowQueryParserNUMBER)
}

func (s *InexpressionmetricContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, i)
}

func (s *InexpressionmetricContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InexpressionmetricContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InexpressionmetricContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterInexpressionmetric(s)
	}
}

func (s *InexpressionmetricContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitInexpressionmetric(s)
	}
}

func (s *InexpressionmetricContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitInexpressionmetric(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Inexpressionmetric() (localctx IInexpressionmetricContext) {
	localctx = NewInexpressionmetricContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BitflowQueryParserRULE_inexpressionmetric)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(312)
		p.Match(BitflowQueryParserINWORD)
	}
	{
		p.SetState(313)
		p.Match(BitflowQueryParserT__11)
	}
	{
		p.SetState(314)
		p.Match(BitflowQueryParserNUMBER)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BitflowQueryParserT__1 {
		{
			p.SetState(315)
			p.Match(BitflowQueryParserT__1)
		}
		{
			p.SetState(316)
			p.Match(BitflowQueryParserNUMBER)
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(322)
		p.Match(BitflowQueryParserT__12)
	}

	return localctx
}

// IHastagContext is an interface to support dynamic dispatch.
type IHastagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHastagContext differentiates from other interfaces.
	IsHastagContext()
}

type HastagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHastagContext() *HastagContext {
	var p = new(HastagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_hastag
	return p
}

func (*HastagContext) IsHastagContext() {}

func NewHastagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HastagContext {
	var p = new(HastagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_hastag

	return p
}

func (s *HastagContext) GetParser() antlr.Parser { return s.parser }

func (s *HastagContext) TAGWORD() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserTAGWORD, 0)
}

func (s *HastagContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BitflowQueryParserSTRING)
}

func (s *HastagContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, i)
}

func (s *HastagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HastagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HastagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterHastag(s)
	}
}

func (s *HastagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitHastag(s)
	}
}

func (s *HastagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitHastag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Hastag() (localctx IHastagContext) {
	localctx = NewHastagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BitflowQueryParserRULE_hastag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(BitflowQueryParserTAGWORD)
	}
	{
		p.SetState(325)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(326)
		p.Match(BitflowQueryParserT__13)
	}
	{
		p.SetState(327)
		p.Match(BitflowQueryParserSTRING)
	}
	{
		p.SetState(328)
		p.Match(BitflowQueryParserT__14)
	}

	return localctx
}

// INotnodeContext is an interface to support dynamic dispatch.
type INotnodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotnodeContext differentiates from other interfaces.
	IsNotnodeContext()
}

type NotnodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotnodeContext() *NotnodeContext {
	var p = new(NotnodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_notnode
	return p
}

func (*NotnodeContext) IsNotnodeContext() {}

func NewNotnodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotnodeContext {
	var p = new(NotnodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_notnode

	return p
}

func (s *NotnodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NotnodeContext) NOT() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNOT, 0)
}

func (s *NotnodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotnodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotnodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterNotnode(s)
	}
}

func (s *NotnodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitNotnode(s)
	}
}

func (s *NotnodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitNotnode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Notnode() (localctx INotnodeContext) {
	localctx = NewNotnodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BitflowQueryParserRULE_notnode)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(BitflowQueryParserNOT)
	}

	return localctx
}

// IEndnodeContext is an interface to support dynamic dispatch.
type IEndnodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndnodeContext differentiates from other interfaces.
	IsEndnodeContext()
}

type EndnodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndnodeContext() *EndnodeContext {
	var p = new(EndnodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_endnode
	return p
}

func (*EndnodeContext) IsEndnodeContext() {}

func NewEndnodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndnodeContext {
	var p = new(EndnodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_endnode

	return p
}

func (s *EndnodeContext) GetParser() antlr.Parser { return s.parser }

func (s *EndnodeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserIDENTIFIER, 0)
}

func (s *EndnodeContext) DECITIMES() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserDECITIMES, 0)
}

func (s *EndnodeContext) STRING() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, 0)
}

func (s *EndnodeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserNUMBER, 0)
}

func (s *EndnodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndnodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndnodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterEndnode(s)
	}
}

func (s *EndnodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitEndnode(s)
	}
}

func (s *EndnodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitEndnode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Endnode() (localctx IEndnodeContext) {
	localctx = NewEndnodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BitflowQueryParserRULE_endnode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(BitflowQueryParserNUMBER-48))|(1<<(BitflowQueryParserSTRING-48))|(1<<(BitflowQueryParserDECITIMES-48))|(1<<(BitflowQueryParserIDENTIFIER-48)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserGT, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserGE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserLT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserLE, 0)
}

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserEQ, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BitflowQueryParserRULE_comparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(BitflowQueryParserGT-43))|(1<<(BitflowQueryParserGE-43))|(1<<(BitflowQueryParserLT-43))|(1<<(BitflowQueryParserLE-43))|(1<<(BitflowQueryParserEQ-43)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserOR, 0)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BitflowQueryParserRULE_binary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowQueryParserAND || _la == BitflowQueryParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathematicalOperationContext is an interface to support dynamic dispatch.
type IMathematicalOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathematicalOperationContext differentiates from other interfaces.
	IsMathematicalOperationContext()
}

type MathematicalOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathematicalOperationContext() *MathematicalOperationContext {
	var p = new(MathematicalOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_mathematicalOperation
	return p
}

func (*MathematicalOperationContext) IsMathematicalOperationContext() {}

func NewMathematicalOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathematicalOperationContext {
	var p = new(MathematicalOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_mathematicalOperation

	return p
}

func (s *MathematicalOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *MathematicalOperationContext) TIMES() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserTIMES, 0)
}

func (s *MathematicalOperationContext) DIVIDED() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserDIVIDED, 0)
}

func (s *MathematicalOperationContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserPLUS, 0)
}

func (s *MathematicalOperationContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserMINUS, 0)
}

func (s *MathematicalOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathematicalOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathematicalOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterMathematicalOperation(s)
	}
}

func (s *MathematicalOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitMathematicalOperation(s)
	}
}

func (s *MathematicalOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitMathematicalOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) MathematicalOperation() (localctx IMathematicalOperationContext) {
	localctx = NewMathematicalOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BitflowQueryParserRULE_mathematicalOperation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BitflowQueryParserPLUS-32))|(1<<(BitflowQueryParserMINUS-32))|(1<<(BitflowQueryParserTIMES-32))|(1<<(BitflowQueryParserDIVIDED-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBoolTokenContext is an interface to support dynamic dispatch.
type IBoolTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTokenContext differentiates from other interfaces.
	IsBoolTokenContext()
}

type BoolTokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTokenContext() *BoolTokenContext {
	var p = new(BoolTokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_boolToken
	return p
}

func (*BoolTokenContext) IsBoolTokenContext() {}

func NewBoolTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTokenContext {
	var p = new(BoolTokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_boolToken

	return p
}

func (s *BoolTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTokenContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserTRUE, 0)
}

func (s *BoolTokenContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserFALSE, 0)
}

func (s *BoolTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterBoolToken(s)
	}
}

func (s *BoolTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitBoolToken(s)
	}
}

func (s *BoolTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitBoolToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) BoolToken() (localctx IBoolTokenContext) {
	localctx = NewBoolTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BitflowQueryParserRULE_boolToken)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BitflowQueryParserTRUE || _la == BitflowQueryParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BitflowQueryParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BitflowQueryParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BitflowQueryParserSTRING)
}

func (s *ListContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BitflowQueryParserSTRING, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BitflowQueryListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BitflowQueryVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BitflowQueryParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BitflowQueryParserRULE_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(342)
				p.Match(BitflowQueryParserSTRING)
			}
			{
				p.SetState(343)
				p.Match(BitflowQueryParserT__1)
			}

		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	{
		p.SetState(349)
		p.Match(BitflowQueryParserSTRING)
	}

	return localctx
}

func (p *BitflowQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *MathematicalSelectionContext = nil
		if localctx != nil {
			t = localctx.(*MathematicalSelectionContext)
		}
		return p.MathematicalSelection_Sempred(t, predIndex)

	case 30:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BitflowQueryParser) MathematicalSelection_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *BitflowQueryParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
