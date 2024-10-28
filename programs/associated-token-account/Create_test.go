package associatedtokenaccount

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func Test_create(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta.RPC)
	wallet := solana.MustPrivateKeyFromBase58("")

	create := NewCreateInstruction(
		wallet.PublicKey(),
		wallet.PublicKey(),
		solana.MustPublicKeyFromBase58("4H2kQgC4hdt35AryM4NjB6RoNuuX4WR6fRMQ4cqcg1zr"),
		solana.TokenProgramID,
	).Build()
	createIdempotent := NewCreateIdempotentInstruction(
		wallet.PublicKey(),
		wallet.PublicKey(),
		solana.MustPublicKeyFromBase58("4H2kQgC4hdt35AryM4NjB6RoNuuX4WR6fRMQ4cqcg1zr"),
		solana.TokenProgramID,
	).Build()

	spew.Dump(create, createIdempotent)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			create,
			createIdempotent,
		},
		solana.Hash{},
		solana.TransactionPayer(wallet.PublicKey()),
	)
	if err != nil {
		t.Fatal(err)
	}

	tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key == wallet.PublicKey() {
			return &wallet
		}
		return nil
	})

	sign, err := client.SimulateTransactionWithOpts(context.Background(), tx, &rpc.SimulateTransactionOpts{ReplaceRecentBlockhash: true})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spew.Sdump(sign))
}

func Test_create2022(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta.RPC)
	wallet := solana.MustPrivateKeyFromBase58("")

	create := NewCreateInstruction(
		wallet.PublicKey(),
		wallet.PublicKey(),
		solana.MustPublicKeyFromBase58("HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC"),
		solana.Token2022ProgramID,
	)
	createIdemptent := NewCreateIdempotentInstruction(
		wallet.PublicKey(),
		wallet.PublicKey(),
		solana.MustPublicKeyFromBase58("HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC"),
		solana.Token2022ProgramID,
	)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			create.Build(),
			createIdemptent.Build(),
		},
		solana.Hash{},
		solana.TransactionPayer(wallet.PublicKey()),
	)
	if err != nil {
		t.Fatal(err)
	}

	tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key == wallet.PublicKey() {
			return &wallet
		}
		return nil
	})

	sign, err := client.SimulateTransactionWithOpts(context.Background(), tx, &rpc.SimulateTransactionOpts{ReplaceRecentBlockhash: true})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spew.Sdump(sign))
}
