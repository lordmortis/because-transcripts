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

func testUtteranceFragments(t *testing.T) {
	t.Parallel()

	query := UtteranceFragments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUtteranceFragmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
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

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtteranceFragmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UtteranceFragments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtteranceFragmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UtteranceFragmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUtteranceFragmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UtteranceFragmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UtteranceFragment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UtteranceFragmentExists to return true, but got false.")
	}
}

func testUtteranceFragmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	utteranceFragmentFound, err := FindUtteranceFragment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if utteranceFragmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUtteranceFragmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UtteranceFragments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUtteranceFragmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UtteranceFragments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUtteranceFragmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	utteranceFragmentOne := &UtteranceFragment{}
	utteranceFragmentTwo := &UtteranceFragment{}
	if err = randomize.Struct(seed, utteranceFragmentOne, utteranceFragmentDBTypes, false, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}
	if err = randomize.Struct(seed, utteranceFragmentTwo, utteranceFragmentDBTypes, false, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = utteranceFragmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = utteranceFragmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UtteranceFragments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUtteranceFragmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	utteranceFragmentOne := &UtteranceFragment{}
	utteranceFragmentTwo := &UtteranceFragment{}
	if err = randomize.Struct(seed, utteranceFragmentOne, utteranceFragmentDBTypes, false, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}
	if err = randomize.Struct(seed, utteranceFragmentTwo, utteranceFragmentDBTypes, false, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = utteranceFragmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = utteranceFragmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func utteranceFragmentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func utteranceFragmentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UtteranceFragment) error {
	*o = UtteranceFragment{}
	return nil
}

func testUtteranceFragmentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UtteranceFragment{}
	o := &UtteranceFragment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment object: %s", err)
	}

	AddUtteranceFragmentHook(boil.BeforeInsertHook, utteranceFragmentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentBeforeInsertHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.AfterInsertHook, utteranceFragmentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentAfterInsertHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.AfterSelectHook, utteranceFragmentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentAfterSelectHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.BeforeUpdateHook, utteranceFragmentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentBeforeUpdateHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.AfterUpdateHook, utteranceFragmentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentAfterUpdateHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.BeforeDeleteHook, utteranceFragmentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentBeforeDeleteHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.AfterDeleteHook, utteranceFragmentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentAfterDeleteHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.BeforeUpsertHook, utteranceFragmentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentBeforeUpsertHooks = []UtteranceFragmentHook{}

	AddUtteranceFragmentHook(boil.AfterUpsertHook, utteranceFragmentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	utteranceFragmentAfterUpsertHooks = []UtteranceFragmentHook{}
}

func testUtteranceFragmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUtteranceFragmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(utteranceFragmentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUtteranceFragmentToManyUtteranceFragmentLinks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UtteranceFragment
	var b, c UtteranceFragmentLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
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

	queries.Assign(&b.UtteranceFragmentID, a.ID)
	queries.Assign(&c.UtteranceFragmentID, a.ID)
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
		if queries.Equal(v.UtteranceFragmentID, b.UtteranceFragmentID) {
			bFound = true
		}
		if queries.Equal(v.UtteranceFragmentID, c.UtteranceFragmentID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UtteranceFragmentSlice{&a}
	if err = a.L.LoadUtteranceFragmentLinks(ctx, tx, false, (*[]*UtteranceFragment)(&slice), nil); err != nil {
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

func testUtteranceFragmentToManyAddOpUtteranceFragmentLinks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UtteranceFragment
	var b, c, d, e UtteranceFragmentLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, utteranceFragmentDBTypes, false, strmangle.SetComplement(utteranceFragmentPrimaryKeyColumns, utteranceFragmentColumnsWithoutDefault)...); err != nil {
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

		if !queries.Equal(a.ID, first.UtteranceFragmentID) {
			t.Error("foreign key was wrong value", a.ID, first.UtteranceFragmentID)
		}
		if !queries.Equal(a.ID, second.UtteranceFragmentID) {
			t.Error("foreign key was wrong value", a.ID, second.UtteranceFragmentID)
		}

		if first.R.UtteranceFragment != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.UtteranceFragment != &a {
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

func testUtteranceFragmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
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

func testUtteranceFragmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UtteranceFragmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUtteranceFragmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UtteranceFragments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	utteranceFragmentDBTypes = map[string]string{`ID`: `BLOB`, `Value`: `TEXT`, `CreatedAt`: `DATETIME`, `UpdatedAt`: `DATETIME`}
	_                        = bytes.MinRead
)

func testUtteranceFragmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(utteranceFragmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(utteranceFragmentAllColumns) == len(utteranceFragmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUtteranceFragmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(utteranceFragmentAllColumns) == len(utteranceFragmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UtteranceFragment{}
	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, utteranceFragmentDBTypes, true, utteranceFragmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(utteranceFragmentAllColumns, utteranceFragmentPrimaryKeyColumns) {
		fields = utteranceFragmentAllColumns
	} else {
		fields = strmangle.SetComplement(
			utteranceFragmentAllColumns,
			utteranceFragmentPrimaryKeyColumns,
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

	slice := UtteranceFragmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUtteranceFragmentsUpsert(t *testing.T) {
	t.Parallel()
	if len(utteranceFragmentAllColumns) == len(utteranceFragmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UtteranceFragment{}
	if err = randomize.Struct(seed, &o, utteranceFragmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UtteranceFragment: %s", err)
	}

	count, err := UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, utteranceFragmentDBTypes, false, utteranceFragmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UtteranceFragment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UtteranceFragment: %s", err)
	}

	count, err = UtteranceFragments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}