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

func testSpeakers(t *testing.T) {
	t.Parallel()

	query := Speakers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpeakersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
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

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpeakersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Speakers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpeakersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpeakerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpeakersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpeakerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Speaker exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpeakerExists to return true, but got false.")
	}
}

func testSpeakersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	speakerFound, err := FindSpeaker(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if speakerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpeakersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Speakers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpeakersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Speakers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpeakersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	speakerOne := &Speaker{}
	speakerTwo := &Speaker{}
	if err = randomize.Struct(seed, speakerOne, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}
	if err = randomize.Struct(seed, speakerTwo, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = speakerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = speakerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Speakers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpeakersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	speakerOne := &Speaker{}
	speakerTwo := &Speaker{}
	if err = randomize.Struct(seed, speakerOne, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}
	if err = randomize.Struct(seed, speakerTwo, speakerDBTypes, false, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = speakerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = speakerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func speakerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func speakerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Speaker) error {
	*o = Speaker{}
	return nil
}

func testSpeakersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Speaker{}
	o := &Speaker{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, speakerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Speaker object: %s", err)
	}

	AddSpeakerHook(boil.BeforeInsertHook, speakerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	speakerBeforeInsertHooks = []SpeakerHook{}

	AddSpeakerHook(boil.AfterInsertHook, speakerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	speakerAfterInsertHooks = []SpeakerHook{}

	AddSpeakerHook(boil.AfterSelectHook, speakerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	speakerAfterSelectHooks = []SpeakerHook{}

	AddSpeakerHook(boil.BeforeUpdateHook, speakerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	speakerBeforeUpdateHooks = []SpeakerHook{}

	AddSpeakerHook(boil.AfterUpdateHook, speakerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	speakerAfterUpdateHooks = []SpeakerHook{}

	AddSpeakerHook(boil.BeforeDeleteHook, speakerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	speakerBeforeDeleteHooks = []SpeakerHook{}

	AddSpeakerHook(boil.AfterDeleteHook, speakerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	speakerAfterDeleteHooks = []SpeakerHook{}

	AddSpeakerHook(boil.BeforeUpsertHook, speakerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	speakerBeforeUpsertHooks = []SpeakerHook{}

	AddSpeakerHook(boil.AfterUpsertHook, speakerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	speakerAfterUpsertHooks = []SpeakerHook{}
}

func testSpeakersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpeakersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(speakerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpeakerToManyUtterances(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Speaker
	var b, c Utterance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, utteranceDBTypes, false, utteranceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SpeakerID, a.ID)
	queries.Assign(&c.SpeakerID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Utterances().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SpeakerID, b.SpeakerID) {
			bFound = true
		}
		if queries.Equal(v.SpeakerID, c.SpeakerID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SpeakerSlice{&a}
	if err = a.L.LoadUtterances(ctx, tx, false, (*[]*Speaker)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Utterances); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Utterances = nil
	if err = a.L.LoadUtterances(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Utterances); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSpeakerToManyAddOpUtterances(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Speaker
	var b, c, d, e Utterance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, speakerDBTypes, false, strmangle.SetComplement(speakerPrimaryKeyColumns, speakerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Utterance{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, utteranceDBTypes, false, strmangle.SetComplement(utterancePrimaryKeyColumns, utteranceColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Utterance{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUtterances(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SpeakerID) {
			t.Error("foreign key was wrong value", a.ID, first.SpeakerID)
		}
		if !queries.Equal(a.ID, second.SpeakerID) {
			t.Error("foreign key was wrong value", a.ID, second.SpeakerID)
		}

		if first.R.Speaker != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Speaker != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Utterances[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Utterances[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Utterances().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSpeakersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
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

func testSpeakersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpeakerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpeakersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Speakers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	speakerDBTypes = map[string]string{`ID`: `BLOB`, `TranscriptName`: `TEXT`, `Name`: `TEXT`, `CreatedAt`: `DATETIME`, `UpdatedAt`: `DATETIME`}
	_              = bytes.MinRead
)

func testSpeakersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(speakerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(speakerAllColumns) == len(speakerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpeakersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(speakerAllColumns) == len(speakerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Speaker{}
	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, speakerDBTypes, true, speakerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(speakerAllColumns, speakerPrimaryKeyColumns) {
		fields = speakerAllColumns
	} else {
		fields = strmangle.SetComplement(
			speakerAllColumns,
			speakerPrimaryKeyColumns,
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

	slice := SpeakerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpeakersUpsert(t *testing.T) {
	t.Parallel()
	if len(speakerAllColumns) == len(speakerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Speaker{}
	if err = randomize.Struct(seed, &o, speakerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Speaker: %s", err)
	}

	count, err := Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, speakerDBTypes, false, speakerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Speaker struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Speaker: %s", err)
	}

	count, err = Speakers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
