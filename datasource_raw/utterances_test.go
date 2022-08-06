// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datasource_raw

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUtterances(t *testing.T) {
	t.Parallel()

	query := Utterances()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUtterancesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtterancesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Utterances().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtterancesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UtteranceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtterancesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UtteranceExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Utterance exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UtteranceExists to return true, but got false.")
	}
}

func testUtterancesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	utteranceFound, err := FindUtterance(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if utteranceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUtterancesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Utterances().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUtterancesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Utterances().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUtterancesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	utteranceOne := &Utterance{}
	utteranceTwo := &Utterance{}
	if err = randomize.Struct(seed, utteranceOne, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}
	if err = randomize.Struct(seed, utteranceTwo, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = utteranceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = utteranceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Utterances().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUtterancesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	utteranceOne := &Utterance{}
	utteranceTwo := &Utterance{}
	if err = randomize.Struct(seed, utteranceOne, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}
	if err = randomize.Struct(seed, utteranceTwo, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = utteranceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = utteranceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func utteranceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func utteranceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Utterance) error {
	*o = Utterance{}
	return nil
}

func testUtterancesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Utterance{}
	o := &Utterance{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, utteranceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Utterance object: %s", err)
	}

	AddUtteranceHook(boil.BeforeInsertHook, utteranceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	utteranceBeforeInsertHooks = []UtteranceHook{}

	AddUtteranceHook(boil.AfterInsertHook, utteranceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	utteranceAfterInsertHooks = []UtteranceHook{}

	AddUtteranceHook(boil.AfterSelectHook, utteranceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	utteranceAfterSelectHooks = []UtteranceHook{}

	AddUtteranceHook(boil.BeforeUpdateHook, utteranceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	utteranceBeforeUpdateHooks = []UtteranceHook{}

	AddUtteranceHook(boil.AfterUpdateHook, utteranceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	utteranceAfterUpdateHooks = []UtteranceHook{}

	AddUtteranceHook(boil.BeforeDeleteHook, utteranceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	utteranceBeforeDeleteHooks = []UtteranceHook{}

	AddUtteranceHook(boil.AfterDeleteHook, utteranceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	utteranceAfterDeleteHooks = []UtteranceHook{}

	AddUtteranceHook(boil.BeforeUpsertHook, utteranceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	utteranceBeforeUpsertHooks = []UtteranceHook{}

	AddUtteranceHook(boil.AfterUpsertHook, utteranceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	utteranceAfterUpsertHooks = []UtteranceHook{}
}

func testUtterancesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUtterancesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(utteranceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUtteranceToManyUtteranceFragmentLinks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c UtteranceFragmentLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, utteranceFragmentLinkDBTypes, false, utteranceFragmentLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, utteranceFragmentLinkDBTypes, false, utteranceFragmentLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UtteranceID, a.ID)
	queries.Assign(&c.UtteranceID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UtteranceFragmentLinks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UtteranceID, b.UtteranceID) {
			bFound = true
		}
		if queries.Equal(v.UtteranceID, c.UtteranceID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UtteranceSlice{&a}
	if err = a.L.LoadUtteranceFragmentLinks(ctx, tx, false, (*[]*Utterance)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UtteranceFragmentLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UtteranceFragmentLinks = nil
	if err = a.L.LoadUtteranceFragmentLinks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UtteranceFragmentLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUtteranceToManySpeakers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c Speaker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"utterance_speakers\" (\"utterance_id\", \"speaker_id\") values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"utterance_speakers\" (\"utterance_id\", \"speaker_id\") values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Speakers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ID, b.ID) {
			bFound = true
		}
		if queries.Equal(v.ID, c.ID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UtteranceSlice{&a}
	if err = a.L.LoadSpeakers(ctx, tx, false, (*[]*Utterance)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Speakers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Speakers = nil
	if err = a.L.LoadSpeakers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Speakers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUtteranceToManyAddOpUtteranceFragmentLinks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c, d, e UtteranceFragmentLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UtteranceFragmentLink{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, utteranceFragmentLinkDBTypes, false, strmangle.SetComplement(utteranceFragmentLinkPrimaryKeyColumns, utteranceFragmentLinkColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UtteranceFragmentLink{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUtteranceFragmentLinks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UtteranceID) {
			t.Error("foreign key was wrong value", a.ID, first.UtteranceID)
		}
		if !queries.Equal(a.ID, second.UtteranceID) {
			t.Error("foreign key was wrong value", a.ID, second.UtteranceID)
		}

		if first.R.Utterance != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Utterance != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UtteranceFragmentLinks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UtteranceFragmentLinks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UtteranceFragmentLinks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testUtteranceToManyAddOpSpeakers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c, d, e Speaker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Speaker{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, speakerDBTypes, false, strmangle.SetComplement(speakerPrimaryKeyColumns, speakerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Speaker{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSpeakers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Utterances[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Utterances[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Speakers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Speakers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Speakers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testUtteranceToManySetOpSpeakers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c, d, e Speaker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Speaker{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, speakerDBTypes, false, strmangle.SetComplement(speakerPrimaryKeyColumns, speakerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSpeakers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Speakers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSpeakers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Speakers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Utterances) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Utterances) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Utterances[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Utterances[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Speakers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Speakers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testUtteranceToManyRemoveOpSpeakers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c, d, e Speaker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Speaker{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, speakerDBTypes, false, strmangle.SetComplement(speakerPrimaryKeyColumns, speakerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSpeakers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Speakers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSpeakers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Speakers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Utterances) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Utterances) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Utterances[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Utterances[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Speakers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Speakers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Speakers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testUtteranceToOneEpisodeUsingEpisode(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Utterance
	var foreign Episode

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, episodeDBTypes, false, episodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Episode struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.EpisodeID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Episode().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UtteranceSlice{&local}
	if err = local.L.LoadEpisode(ctx, tx, false, (*[]*Utterance)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Episode == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Episode = nil
	if err = local.L.LoadEpisode(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Episode == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUtteranceToOneSetOpEpisodeUsingEpisode(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Utterance
	var b, c Episode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, episodeDBTypes, false, strmangle.SetComplement(episodePrimaryKeyColumns, episodeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, episodeDBTypes, false, strmangle.SetComplement(episodePrimaryKeyColumns, episodeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Episode{&b, &c} {
		err = a.SetEpisode(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Episode != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Utterances[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.EpisodeID, x.ID) {
			t.Error("foreign key was wrong value", a.EpisodeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EpisodeID))
		reflect.Indirect(reflect.ValueOf(&a.EpisodeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.EpisodeID, x.ID) {
			t.Error("foreign key was wrong value", a.EpisodeID, x.ID)
		}
	}
}

func testUtterancesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUtterancesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UtteranceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUtterancesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Utterances().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	utteranceDBTypes = map[string]string{`ID`: `BLOB`, `EpisodeID`: `BLOB`, `SequenceNo`: `INTEGER`, `IsParalinguistic`: `INTEGER`, `StartTime`: `INTEGER`, `EndTime`: `INTEGER`, `Utterance`: `TEXT`, `CreatedAt`: `DATETIME`, `UpdatedAt`: `DATETIME`}
	_                = bytes.MinRead
)

func testUtterancesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(utterancePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(utteranceAllColumns) == len(utterancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utterancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUtterancesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(utteranceAllColumns) == len(utterancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Utterance{}
	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utteranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, utteranceDBTypes, true, utterancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(utteranceAllColumns, utterancePrimaryKeyColumns) {
		fields = utteranceAllColumns
	} else {
		fields = strmangle.SetComplement(
			utteranceAllColumns,
			utterancePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UtteranceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUtterancesUpsert(t *testing.T) {
	t.Parallel()
	if len(utteranceAllColumns) == len(utterancePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Utterance{}
	if err = randomize.Struct(seed, &o, utteranceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Utterance: %s", err)
	}

	count, err := Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, utteranceDBTypes, false, utterancePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Utterance struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Utterance: %s", err)
	}

	count, err = Utterances().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
