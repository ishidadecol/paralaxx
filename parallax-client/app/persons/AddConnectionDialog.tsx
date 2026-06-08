"use client";

import { useState } from "react";
import { Plus } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";

import { Field, FieldGroup } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useEntityConnection } from "@/hooks/useEntityConnection";
import { usePersons } from "@/hooks/usePersons";
import { useParams } from "next/navigation";

type EntityType = "person";

const relationshipOptions: Record<EntityType, string[]> = {
  person: ["KNOWS", "WORKED_WITH", "FAMILY_WITH"],
};

export function AddConnectionDialog() {
  //Connection hook
  const { connections, loading, addEntityConnection } = useEntityConnection();

  const { persons } = usePersons();

  const params = useParams();
  const { id } = params;

  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  const [entityType, setEntityType] = useState<EntityType | "">("");

  const [targetEntityId, setTargetEntityId] = useState("");

  const [relationshipType, setRelationshipType] = useState("");

  //Handle submit
  const handleSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    console.log({
        sourceEntityId: id,
        targetEntityId,
        relationshipType,
        startDate,
        endDate,
      });
    try {
      await addEntityConnection({
        sourceEntityId: id as string,
        targetEntityId,
        relationshipType,
        startDate: startDate || null,
        endDate: endDate || null,
      });
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">
          <Plus className="mr-2 h-4 w-4" />
          Add Connection
        </Button>
      </DialogTrigger>

      <DialogContent>
        <form onSubmit={handleSubmit} className="space-y-4">
          <DialogHeader>
            <DialogTitle>Add Connection</DialogTitle>

            <DialogDescription>
              Create a relationship between this entity and another entity.
            </DialogDescription>
          </DialogHeader>

          <FieldGroup>
            <Field>
              <Label>Entity Type</Label>

              <Select
                value={entityType}
                onValueChange={(value) => {
                  setEntityType(value as EntityType);

                  setTargetEntityId("");
                  setRelationshipType("");

                  // TODO:
                  // Fetch entities by type
                }}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select entity type" />
                </SelectTrigger>

                <SelectContent>
                  <SelectItem value="person">Person</SelectItem>
                </SelectContent>
              </Select>
            </Field>

            <Field>
              <Label>Target Entity</Label>

              <Select
                value={targetEntityId}
                onValueChange={setTargetEntityId}
                disabled={!entityType}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select entity" />
                </SelectTrigger>

                <SelectContent>
                  {persons.map((person) => (
                    <SelectItem key={person.entity_id} value={person.entity_id}>
                      {person.first_name} {person.last_name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </Field>

            <Field>
              <Label>Relationship</Label>

              <Select
                value={relationshipType}
                onValueChange={setRelationshipType}
                disabled={!entityType}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select relationship" />
                </SelectTrigger>

                <SelectContent>
                  {entityType &&
                    relationshipOptions[entityType]?.map((relationship) => (
                      <SelectItem key={relationship} value={relationship}>
                        {relationship}
                      </SelectItem>
                    ))}
                </SelectContent>
              </Select>
            </Field>

            <Field>
              <Label>Start Date</Label>

              <Input
                type="date"
                name="startDate"
                onChange={(e) => setStartDate(e.target.value)}
              />
            </Field>

            <Field>
              <Label>End Date</Label>

              <Input
                type="date"
                name="endDate"
                onChange={(e) => setEndDate(e.target.value)}
              />
            </Field>
          </FieldGroup>

          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>

            <Button type="submit">Create Connection</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
