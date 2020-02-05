// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package reward

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *RewardActorState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.RewardMap (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.RewardMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.RewardMap: %w", err)
	}

	return nil
}

func (t *RewardActorState) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.RewardMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.RewardMap: %w", err)
		}

		t.RewardMap = c

	}
	return nil
}

func (t *AwardBlockRewardParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Penalty (big.Int) (struct)
	if err := t.Penalty.MarshalCBOR(w); err != nil {
		return err
	}

	// t.GasReward (big.Int) (struct)
	if err := t.GasReward.MarshalCBOR(w); err != nil {
		return err
	}

	// t.NominalPower (big.Int) (struct)
	if err := t.NominalPower.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PledgeCollateral (big.Int) (struct)
	if err := t.PledgeCollateral.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AwardBlockRewardParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Penalty (big.Int) (struct)

	{

		if err := t.Penalty.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.GasReward (big.Int) (struct)

	{

		if err := t.GasReward.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.NominalPower (big.Int) (struct)

	{

		if err := t.NominalPower.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.PledgeCollateral (big.Int) (struct)

	{

		if err := t.PledgeCollateral.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *Reward) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.VestingFunction (reward.VestingFunction) (int64)
	if t.VestingFunction >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.VestingFunction))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.VestingFunction)-1)); err != nil {
			return err
		}
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if t.StartEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.StartEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.StartEpoch)-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if t.EndEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.EndEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.EndEpoch)-1)); err != nil {
			return err
		}
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}

	// t.AmountWithdrawn (big.Int) (struct)
	if err := t.AmountWithdrawn.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Reward) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.VestingFunction (reward.VestingFunction) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.VestingFunction = VestingFunction(extraI)
	}
	// t.StartEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.StartEpoch = abi.ChainEpoch(extraI)
	}
	// t.EndEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EndEpoch = abi.ChainEpoch(extraI)
	}
	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.AmountWithdrawn (big.Int) (struct)

	{

		if err := t.AmountWithdrawn.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}