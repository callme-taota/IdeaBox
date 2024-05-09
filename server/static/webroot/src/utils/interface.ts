export interface UserItem {
    ID: number
    Name: string
    Color: string
}

export interface IdeaItem {
    ID: number
    Name: string
    Description: string
    Author: number
    CreateAt: string
    UpdateAt: string
}

export interface FullIdeaItem {
    IdeaID: number
    IdeaName: string
    Description: string
    Author: number
    CreatedAt: string
    UpdatedAt: string
    Color: string
    Name: string
}

export interface Idea {
    Name: string
    Description: string
    Author: number
}

export interface IdeaWithID {
    ID: number
    Name: string
    Description: string
    Author: number
}