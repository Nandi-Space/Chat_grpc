package model

import (
	"errors"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Chat struct {
	ID         string `rethinkdb:"id"`
	SenderID   string `rethinkdb:"sender_id"`
	SenderName string `rethinkdb:"sender_name"`
	InboxHash  string `rethinkdb:"inbox_hash"`
	Msg        string `rethinkdb:"msg"`
	File       string `rethinkdb:"file,omitempty"`
	Meta       string `rethinkdb:"meta,omitempty"`
	DeletedBy  string `rethinkdb:"deleted_by,omitempty"`
	CreatedAt  int64  `gorethink:"createdAt"`
}

type ChatList []*Chat

func (c *Chat) Validate() error {
	if c.SenderID == "" {
		return errors.New("sender id is empty")
	}
	if c.SenderName == "" {
		return errors.New("sender name is empty")
	}
	if c.InboxHash == "" {
		return errors.New("inbox hash id is empty")
	}
	if c.Msg == "" {
		return errors.New("message id is empty")
	}
	return nil
}

func (c *Chat) Insert() error {
	_, err := r.Table("CHAT").Insert(c).RunWrite(handler)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chat) Delete() error {
	if c.ID == "" {
		return errors.New("sender id is empty")
	}
	if c.DeletedBy == "" {
		return errors.New("deleted by is empty")
	}

	_, err := r.Table("CHAT").Filter(&Chat{ID: c.ID}).Update(&Chat{DeletedBy: c.DeletedBy}).RunWrite(handler)
	if err != nil {
		return errors.New("couldn't delete chat")
	}
	return nil
}

func (c *Chat) Fetch(cl *ChatList) error {
	if c.InboxHash == "" {
		return errors.New("room id is empty")
	}

	result, err := r.Table("CHAT").Filter(&Chat{InboxHash: c.InboxHash}).Run(handler)
	if err != nil {
		return err
	}
	result.All(cl)
	return nil
}
