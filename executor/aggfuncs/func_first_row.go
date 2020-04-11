// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package aggfuncs

import (
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/types/json"
	"github.com/pingcap/tidb/util/chunk"
	"github.com/pingcap/tidb/util/stringutil"
)

type basePartialResult4FirstRow struct {
	// isNull indicates whether the first row is null.
	isNull bool
	// gotFirstRow indicates whether the first row has been got,
	// if so, we would avoid evaluating the values of the remained rows.
	gotFirstRow bool
}

type partialResult4FirstRowInt struct {
	basePartialResult4FirstRow

	val int64
}

type partialResult4FirstRowFloat32 struct {
	basePartialResult4FirstRow

	val float32
}

type partialResult4FirstRowDecimal struct {
	basePartialResult4FirstRow

	val types.MyDecimal
}

type partialResult4FirstRowFloat64 struct {
	basePartialResult4FirstRow

	val float64
}

type partialResult4FirstRowString struct {
	basePartialResult4FirstRow

	val string
}

type partialResult4FirstRowTime struct {
	basePartialResult4FirstRow

	val types.Time
}

type partialResult4FirstRowDuration struct {
	basePartialResult4FirstRow

	val types.Duration
}

type partialResult4FirstRowJSON struct {
	basePartialResult4FirstRow

	val json.BinaryJSON
}

type partialResult4FirstRowEnum struct {
	basePartialResult4FirstRow

	val types.Enum
}

type partialResult4FirstRowSet struct {
	basePartialResult4FirstRow

	val types.Set
}

type firstRow4Int struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Int) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowInt))
}

func (e *firstRow4Int) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowInt)(pr)
	p.val, p.isNull, p.gotFirstRow = 0, false, false
}

func (e *firstRow4Int) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowInt)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalInt(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, input
		e.accquiredRows++
	}
	return nil
}

func (*firstRow4Int) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowInt)(src), (*partialResult4FirstRowInt)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4Int) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowInt)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendInt64(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Int) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Float32 struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Float32) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowFloat32))
}

func (e *firstRow4Float32) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowFloat32)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Float32) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowFloat32)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalReal(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, float32(input)
		e.accquiredRows++
	}
	return nil
}
func (*firstRow4Float32) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowFloat32)(src), (*partialResult4FirstRowFloat32)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4Float32) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowFloat32)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendFloat32(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Float32) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Float64 struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Float64) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowFloat64))
}

func (e *firstRow4Float64) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowFloat64)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Float64) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowFloat64)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalReal(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, input
		e.accquiredRows++
	}
	return nil
}

func (*firstRow4Float64) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowFloat64)(src), (*partialResult4FirstRowFloat64)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}
func (e *firstRow4Float64) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowFloat64)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendFloat64(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Float64) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4String struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4String) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowString))
}

func (e *firstRow4String) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowString)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4String) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowString)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalString(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, stringutil.Copy(input)
		e.accquiredRows++
	}
	return nil
}

func (*firstRow4String) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowString)(src), (*partialResult4FirstRowString)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4String) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowString)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendString(e.ordinal, p.val)
	return nil
}

func (e *firstRow4String) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Time struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Time) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowTime))
}

func (e *firstRow4Time) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowTime)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Time) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowTime)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalTime(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, input
		e.accquiredRows++
	}
	return nil
}
func (*firstRow4Time) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowTime)(src), (*partialResult4FirstRowTime)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil

}

func (e *firstRow4Time) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowTime)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendTime(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Time) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Duration struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Duration) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowDuration))
}

func (e *firstRow4Duration) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowDuration)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Duration) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowDuration)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalDuration(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, input
		e.accquiredRows++
	}
	return nil
}
func (*firstRow4Duration) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowDuration)(src), (*partialResult4FirstRowDuration)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4Duration) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowDuration)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendDuration(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Duration) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4JSON struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4JSON) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowJSON))
}

func (e *firstRow4JSON) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowJSON)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4JSON) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowJSON)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalJSON(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, isNull, input.Copy()
		e.accquiredRows++
	}
	return nil
}
func (*firstRow4JSON) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowJSON)(src), (*partialResult4FirstRowJSON)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4JSON) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowJSON)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendJSON(e.ordinal, p.val)
	return nil
}

func (e *firstRow4JSON) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Decimal struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Decimal) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowDecimal))
}

func (e *firstRow4Decimal) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowDecimal)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Decimal) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowDecimal)(pr)
	if p.gotFirstRow {
		return nil
	}
	if len(rowsInGroup) > 0 {
		input, isNull, err := e.args[0].EvalDecimal(sctx, rowsInGroup[0])
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull = true, isNull
		e.accquiredRows++
		if input != nil {
			p.val = *input
		}
	}
	return nil
}

func (e *firstRow4Decimal) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowDecimal)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendMyDecimal(e.ordinal, &p.val)
	return nil
}

func (*firstRow4Decimal) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowDecimal)(src), (*partialResult4FirstRowDecimal)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}

func (e *firstRow4Decimal) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Enum struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Enum) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowEnum))
}

func (e *firstRow4Enum) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowEnum)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Enum) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowEnum)(pr)
	if p.gotFirstRow {
		return nil
	}
	for _, row := range rowsInGroup {
		d, err := e.args[0].Eval(row)
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, d.IsNull(), d.GetMysqlEnum()
		e.accquiredRows++
		break
	}
	return nil
}

func (*firstRow4Enum) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowEnum)(src), (*partialResult4FirstRowEnum)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}
func (e *firstRow4Enum) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowEnum)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendEnum(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Enum) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}

type firstRow4Set struct {
	baseAggFunc
	// accquiredRows is the number of scanned distinct value
	accquiredRows int
}

func (e *firstRow4Set) AllocPartialResult() PartialResult {
	return PartialResult(new(partialResult4FirstRowSet))
}

func (e *firstRow4Set) ResetPartialResult(pr PartialResult) {
	p := (*partialResult4FirstRowSet)(pr)
	p.isNull, p.gotFirstRow = false, false
}

func (e *firstRow4Set) UpdatePartialResult(sctx sessionctx.Context, rowsInGroup []chunk.Row, pr PartialResult) error {
	p := (*partialResult4FirstRowSet)(pr)
	if p.gotFirstRow {
		return nil
	}
	for _, row := range rowsInGroup {
		d, err := e.args[0].Eval(row)
		if err != nil {
			return err
		}
		p.gotFirstRow, p.isNull, p.val = true, d.IsNull(), d.GetMysqlSet()
		e.accquiredRows++
		break
	}
	return nil
}

func (*firstRow4Set) MergePartialResult(sctx sessionctx.Context, src PartialResult, dst PartialResult) error {
	p1, p2 := (*partialResult4FirstRowSet)(src), (*partialResult4FirstRowSet)(dst)
	if !p2.gotFirstRow {
		*p2 = *p1
	}
	return nil
}
func (e *firstRow4Set) AppendFinalResult2Chunk(sctx sessionctx.Context, pr PartialResult, chk *chunk.Chunk) error {
	p := (*partialResult4FirstRowSet)(pr)
	if p.isNull || !p.gotFirstRow {
		chk.AppendNull(e.ordinal)
		return nil
	}
	chk.AppendSet(e.ordinal, p.val)
	return nil
}

func (e *firstRow4Set) AccquiredEnoughRows(requiredRows int) bool {
	return e.accquiredRows >= requiredRows
}
