package blockchain

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Nft struct {
	Superrare  []int
	Rare       []int
	LessCommon []int
	Common     []int
}

const (
	MINING_DIFFICULTY  = 2
	MINING_SENDER      = "METCHAIN_Blockchain"
	MINING_REWARD      = 0.3
	MINING_REWARD_MEGA = 3
	MINING_REWARD_MET  = 15
	MINING_TIMER_SEC   = 10
	STAKING_SENDER     = "Coinbase"
)

func (bc *Blockchain) GetNFT() *Nft {
	nft := new(Nft)
	nft.Superrare = []int{8, 121, 244, 272, 281, 478, 59, 109, 123, 197, 225, 270, 321, 387, 394, 430, 468, 483}
	nft.Rare = []int{6, 69, 98, 120, 140, 141, 147, 150, 168, 186, 195, 201, 251, 288, 332, 419, 457, 465, 469, 496, 24, 47, 103, 157, 158, 175, 203, 218, 228, 286, 295, 314, 356, 361, 368, 371, 376, 386, 409, 488, 29, 102, 110, 116, 156, 163, 174, 180, 206, 283, 310, 323, 339, 346, 360, 366, 395, 454, 462, 487, 37, 60, 142, 145, 204, 221, 238, 260, 269, 280, 370, 375, 408, 410, 435, 464, 467, 479, 497, 96, 128, 137, 152, 164, 188, 196, 248, 282, 294, 305, 335, 338, 340, 347, 351, 389, 428, 481, 493}
	nft.LessCommon = []int{4, 16, 32, 34, 52, 61, 63, 73, 82, 86, 91, 131, 133, 149, 162, 173, 192, 223, 246, 268, 285, 287, 302, 312, 344, 363, 367, 396, 403, 411, 422, 444, 459, 475, 494, 9, 14, 38, 39, 68, 99, 112, 130, 144, 166, 176, 230, 233, 241, 259, 264, 277, 289, 296, 297, 327, 342, 354, 359, 365, 379, 383, 398, 405, 417, 441, 466, 11, 12, 30, 31, 58, 74, 78, 117, 148, 200, 205, 215, 245, 249, 324, 353, 355, 406, 423, 432, 433, 440, 452, 453, 473, 474, 500, 15, 21, 45, 46, 56, 125, 170, 214, 243, 257, 301, 304, 316, 317, 320, 325, 349, 357, 362, 380, 385, 393, 400, 429, 490, 18, 51, 66, 84, 101, 108, 126, 135, 161, 165, 172, 177, 191, 202, 237, 252, 262, 284, 290, 298, 329, 343, 420, 421, 463, 471, 482, 486}
	nft.Common = []int{1, 13, 17, 19, 27, 41, 53, 80, 106, 119, 136, 143, 159, 183, 184, 217, 222, 234, 235, 242, 247, 250, 253, 274, 293, 318, 319, 348, 358, 399, 426, 445, 447, 460, 461, 480, 492, 2, 7, 23, 26, 44, 57, 65, 79, 89, 92, 97, 100, 104, 122, 132, 139, 153, 154, 167, 171, 178, 182, 187, 189, 213, 258, 266, 273, 276, 308, 311, 315, 326, 333, 336, 372, 373, 381, 382, 384, 412, 416, 427, 436, 439, 443, 451, 456, 470, 472, 476, 477, 484, 498, 3, 10, 36, 42, 43, 48, 67, 71, 72, 75, 83, 87, 90, 93, 95, 114, 129, 134, 146, 155, 193, 194, 209, 224, 229, 232, 236, 239, 255, 261, 263, 275, 278, 292, 303, 306, 322, 331, 337, 341, 350, 352, 374, 388, 404, 413, 418, 431, 437, 438, 449, 485, 499, 5, 22, 28, 35, 49, 70, 81, 85, 105, 111, 115, 138, 160, 181, 198, 210, 219, 231, 254, 265, 271, 279, 291, 300, 307, 309, 313, 334, 390, 391, 392, 397, 401, 415, 425, 442, 446, 448, 450, 489, 491, 20, 25, 33, 40, 50, 54, 55, 62, 64, 76, 77, 88, 94, 107, 113, 118, 124, 127, 151, 169, 179, 185, 190, 199, 207, 208, 211, 212, 216, 220, 226, 227, 240, 256, 267, 299, 328, 330, 345, 364, 369, 377, 378, 402, 407, 414, 424, 434, 455, 458, 495}
	return nft
}

func (bc *Blockchain) ValidProof(nonce uint64, previousHash [32]byte, transactions []*Transaction, difficulty int, megablock [32]byte, metblock [32]byte) bool {
	zeros := strings.Repeat("0", difficulty)
	timestamp := time.Now().UnixNano()
	guessBlock := MiniBlock{nonce: nonce, timestamp: timestamp, previousHash: previousHash, transactions: transactions, megablock: megablock, metblock: metblock}

	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	log.Println(guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() uint64 {
	transactions := bc.CopyTransactionPool()
	//previousHash := LastMiniBlock().Hash()
	previousHash := [32]byte{}

	megablock := [32]byte{}
	metblock := [32]byte{}
	nonce := uint64(0)
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY, megablock, metblock) {
		nonce += 1

	}
	time.Sleep(10000)
	return nonce
}
