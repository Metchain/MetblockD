package metwallet

import (
	"strings"
)

func CreateFSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)

	i := 1
	j := 3
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateSSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 2
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateTSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 4
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateFoSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 5
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateFiSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 6
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateSiSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 7
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateSeSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 7
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateEiSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 8
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateNSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 9
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}

func CreateTeSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 1
	j := 10
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}
func CreateElSeed(mnemonic string) string {
	s := " "
	nseed := strings.Split(mnemonic, s)
	i := 11
	j := 7
	oldkeyval := nseed[i]
	newkeyval := nseed[j]
	nseed[j] = oldkeyval
	nseed[i] = newkeyval
	newseed := strings.Join([]string(nseed), " ")
	return newseed
}
