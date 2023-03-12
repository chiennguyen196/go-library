// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testPatrons(t *testing.T) {
	t.Parallel()

	query := Patrons()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPatronsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Patrons().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatronSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Patrons().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatronSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatronsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PatronExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Patron exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PatronExists to return true, but got false.")
	}
}

func testPatronsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	patronFound, err := FindPatron(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if patronFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPatronsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Patrons().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPatronsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Patrons().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPatronsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	patronOne := &Patron{}
	patronTwo := &Patron{}
	if err = randomize.Struct(seed, patronOne, patronDBTypes, false, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}
	if err = randomize.Struct(seed, patronTwo, patronDBTypes, false, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patronOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patronTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patrons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPatronsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	patronOne := &Patron{}
	patronTwo := &Patron{}
	if err = randomize.Struct(seed, patronOne, patronDBTypes, false, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}
	if err = randomize.Struct(seed, patronTwo, patronDBTypes, false, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patronOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patronTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func patronBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func patronAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patron) error {
	*o = Patron{}
	return nil
}

func testPatronsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Patron{}
	o := &Patron{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, patronDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Patron object: %s", err)
	}

	AddPatronHook(boil.BeforeInsertHook, patronBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	patronBeforeInsertHooks = []PatronHook{}

	AddPatronHook(boil.AfterInsertHook, patronAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	patronAfterInsertHooks = []PatronHook{}

	AddPatronHook(boil.AfterSelectHook, patronAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	patronAfterSelectHooks = []PatronHook{}

	AddPatronHook(boil.BeforeUpdateHook, patronBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	patronBeforeUpdateHooks = []PatronHook{}

	AddPatronHook(boil.AfterUpdateHook, patronAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	patronAfterUpdateHooks = []PatronHook{}

	AddPatronHook(boil.BeforeDeleteHook, patronBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	patronBeforeDeleteHooks = []PatronHook{}

	AddPatronHook(boil.AfterDeleteHook, patronAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	patronAfterDeleteHooks = []PatronHook{}

	AddPatronHook(boil.BeforeUpsertHook, patronBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	patronBeforeUpsertHooks = []PatronHook{}

	AddPatronHook(boil.AfterUpsertHook, patronAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	patronAfterUpsertHooks = []PatronHook{}
}

func testPatronsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatronsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(patronColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatronToManyHolds(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patron
	var b, c Hold

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holdDBTypes, false, holdColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PatronID = a.ID
	c.PatronID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Holds().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PatronID == b.PatronID {
			bFound = true
		}
		if v.PatronID == c.PatronID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PatronSlice{&a}
	if err = a.L.LoadHolds(ctx, tx, false, (*[]*Patron)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Holds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Holds = nil
	if err = a.L.LoadHolds(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Holds); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPatronToManyOverdueCheckouts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patron
	var b, c OverdueCheckout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, overdueCheckoutDBTypes, false, overdueCheckoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, overdueCheckoutDBTypes, false, overdueCheckoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PatronID = a.ID
	c.PatronID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OverdueCheckouts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PatronID == b.PatronID {
			bFound = true
		}
		if v.PatronID == c.PatronID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PatronSlice{&a}
	if err = a.L.LoadOverdueCheckouts(ctx, tx, false, (*[]*Patron)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OverdueCheckouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OverdueCheckouts = nil
	if err = a.L.LoadOverdueCheckouts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OverdueCheckouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPatronToManyAddOpHolds(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patron
	var b, c, d, e Hold

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patronDBTypes, false, strmangle.SetComplement(patronPrimaryKeyColumns, patronColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Hold{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, holdDBTypes, false, strmangle.SetComplement(holdPrimaryKeyColumns, holdColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Hold{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddHolds(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PatronID {
			t.Error("foreign key was wrong value", a.ID, first.PatronID)
		}
		if a.ID != second.PatronID {
			t.Error("foreign key was wrong value", a.ID, second.PatronID)
		}

		if first.R.Patron != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Patron != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Holds[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Holds[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Holds().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPatronToManyAddOpOverdueCheckouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patron
	var b, c, d, e OverdueCheckout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patronDBTypes, false, strmangle.SetComplement(patronPrimaryKeyColumns, patronColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OverdueCheckout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, overdueCheckoutDBTypes, false, strmangle.SetComplement(overdueCheckoutPrimaryKeyColumns, overdueCheckoutColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*OverdueCheckout{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOverdueCheckouts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PatronID {
			t.Error("foreign key was wrong value", a.ID, first.PatronID)
		}
		if a.ID != second.PatronID {
			t.Error("foreign key was wrong value", a.ID, second.PatronID)
		}

		if first.R.Patron != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Patron != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OverdueCheckouts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OverdueCheckouts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OverdueCheckouts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPatronsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
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

func testPatronsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatronSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatronsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patrons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	patronDBTypes = map[string]string{`ID`: `character varying`, `PatronType`: `enum.patron_type('Regular','Researcher')`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_             = bytes.MinRead
)

func testPatronsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(patronPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(patronAllColumns) == len(patronPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patronDBTypes, true, patronPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPatronsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(patronAllColumns) == len(patronPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patron{}
	if err = randomize.Struct(seed, o, patronDBTypes, true, patronColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patronDBTypes, true, patronPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(patronAllColumns, patronPrimaryKeyColumns) {
		fields = patronAllColumns
	} else {
		fields = strmangle.SetComplement(
			patronAllColumns,
			patronPrimaryKeyColumns,
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

	slice := PatronSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPatronsUpsert(t *testing.T) {
	t.Parallel()

	if len(patronAllColumns) == len(patronPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Patron{}
	if err = randomize.Struct(seed, &o, patronDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patron: %s", err)
	}

	count, err := Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, patronDBTypes, false, patronPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patron struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patron: %s", err)
	}

	count, err = Patrons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
