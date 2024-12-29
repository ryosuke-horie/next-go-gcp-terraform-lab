"use client";

import React from 'react';
import { ListItem, ListItemText, IconButton, Checkbox } from '@mui/material';
import Button from "@mui/material/Button";
import { TaskResponse } from '../types/Task';

interface TodoItemProps {
  task: TaskResponse;
  onToggle: (id: number) => void;
  onDelete: (id: number) => void;
}

const TodoItem: React.FC<TodoItemProps> = ({ task, onToggle, onDelete }) => {
  const handleToggle = () => {
    onToggle(task.id);
    alert(`Toggled TODO with id: ${task.id}`);
  };

  const handleDelete = () => {
    onDelete(task.id);
    alert(`Deleted TODO with id: ${task.id}`);
  };

  return (
    <ListItem
      secondaryAction={
        <IconButton edge="end" aria-label="delete" onClick={handleDelete}>
          <Button variant='contained'>削除</Button>
        </IconButton>
      }
    >
      <Checkbox
        edge="start"
        checked={task.is_completed}
        tabIndex={-1}
        disableRipple
        onChange={handleToggle}
      />
      <ListItemText 
        primary={task.title} 
        secondary={task.detail} 
        style={{ textDecoration: task.is_completed ? 'line-through' : 'none' }} 
      />
    </ListItem>
  );
};

export default TodoItem;
