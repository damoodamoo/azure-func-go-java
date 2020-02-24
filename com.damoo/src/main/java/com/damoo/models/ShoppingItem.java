package com.damoo.models;

public class ShoppingItem {

    private int Id;
    private String Name;

    public String getName() {
        return Name;
    }

    public int getId() {
        return Id;
    }

    public void setId(int id) {
        this.Id = id;
    }

    public void setName(String name) {
        this.Name = name;
    }
}