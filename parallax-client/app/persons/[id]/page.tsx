"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import { Button } from "@/components/ui/button";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import { Badge } from "@/components/ui/badge";

import { Pencil, Plus, Building2, Users, FileText, Mail, Globe, User, Phone } from "lucide-react";

import Image from "next/image";

export default function PersonDetailsPage() {
  // Replace with API data later
  const person = {
    id: "1",
    first_name: "Daniel",
    last_name: "Ishida",
    birth_date: "2003-10-30",
    gender: "Male",
    created_at: "2026-05-01",
    updated_at: "2026-05-30",
  };

  return (
    <div className="container mx-auto p-6 space-y-6">
      {/* HEADER */}
      <Card>
        <CardContent className="flex items-start justify-between py-6">
          <div className="flex gap-5">
            <Image
              src={
                person.gender === "Female"
                  ? "/placeholder_female.jpg"
                  : "/placeholder_male.jpg"
              }
              alt={`${person.first_name} ${person.last_name}`}
              className="w-24 h-24 object-cover"
              width={96}
              height={96}
            />

            <div>
              <h1 className="text-3xl font-bold">
                {person.first_name} {person.last_name}
              </h1>

              <div className="flex gap-2">
                <div className="mt-2 flex gap-2">
                  <Badge>{person.gender}</Badge>
                </div>
                <div className="mt-2 flex gap-2">
                  <Badge>12 connections</Badge>
                </div>
              </div>

              <div className="flex gap-2 mt-5">
                <p className="text-xsm text-muted-foreground">Created At</p>
                <p>{person.created_at}</p>
                <p className="text-xsm text-muted-foreground">Updated At</p>
                <p>{person.updated_at}</p>
              </div>
            </div>
          </div>

          <div className="flex gap-2">
            <Button>
              <Pencil className="mr-2 h-4 w-4" />
              Edit Person
            </Button>

            <Button variant="outline">
              <Plus className="mr-2 h-4 w-4" />
              Add Note
            </Button>

            <Button variant="outline">
              <Plus className="mr-2 h-4 w-4" />
              Add Connection
            </Button>
          </div>
        </CardContent>
      </Card>

      {/* MAIN CONTENT */}
      <Tabs defaultValue="connections">
        <TabsList>
        <TabsTrigger value="Truth Layer">
            Truth layer
        </TabsTrigger>
          <TabsTrigger value="connections">
            <Users className="mr-2 h-4 w-4" />
            Connections
          </TabsTrigger>

          <TabsTrigger value="companies">
            <Building2 className="mr-2 h-4 w-4" />
            Companies
          </TabsTrigger>

          <TabsTrigger value="notes">
            <FileText className="mr-2 h-4 w-4" />
            Notes
          </TabsTrigger>
          <TabsTrigger value="emails">
            <Mail className="mr-2 h-4 w-4" />
            Emails
          </TabsTrigger>
          <TabsTrigger value="phones">
            <Phone className="mr-2 h-4 w-4" />
            Phones
          </TabsTrigger>
          <TabsTrigger value="social">
            <User className="mr-2 h-4 w-4" />
            Social Accounts
          </TabsTrigger>
          <TabsTrigger value="Locations">
            <Globe className="mr-2 h-4 w-4" />
            location
          </TabsTrigger>
        </TabsList>

        {/* CONNECTIONS */}
        <TabsContent value="connections">
          <Card>
            <CardHeader>
              <CardTitle>Connections</CardTitle>
              <CardDescription>People related to this person</CardDescription>
            </CardHeader>

            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Person</TableHead>
                    <TableHead>Relationship</TableHead>
                    <TableHead>Confidence</TableHead>
                  </TableRow>
                </TableHeader>

                <TableBody>
                  <TableRow>
                    <TableCell>Fernanda</TableCell>
                    <TableCell>
                      <Badge>DATED</Badge>
                    </TableCell>
                    <TableCell>High</TableCell>
                  </TableRow>

                  <TableRow>
                    <TableCell>João</TableCell>
                    <TableCell>
                      <Badge>FRIEND</Badge>
                    </TableCell>
                    <TableCell>Medium</TableCell>
                  </TableRow>

                  <TableRow>
                    <TableCell>Pedro</TableCell>
                    <TableCell>
                      <Badge>WORKED_WITH</Badge>
                    </TableCell>
                    <TableCell>High</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </CardContent>
          </Card>
        </TabsContent>

        {/* COMPANIES */}
        <TabsContent value="companies">
          <Card>
            <CardHeader>
              <CardTitle>Companies</CardTitle>
              <CardDescription>Known employment history</CardDescription>
            </CardHeader>

            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Company</TableHead>
                    <TableHead>Role</TableHead>
                    <TableHead>Period</TableHead>
                  </TableRow>
                </TableHeader>

                <TableBody>
                  <TableRow>
                    <TableCell>Lumi Health</TableCell>
                    <TableCell>Founder</TableCell>
                    <TableCell>2024 - Present</TableCell>
                  </TableRow>

                  <TableRow>
                    <TableCell>Acme Corp</TableCell>
                    <TableCell>Software Engineer</TableCell>
                    <TableCell>2022 - 2024</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </CardContent>
          </Card>
        </TabsContent>

        {/* NOTES */}
        <TabsContent value="notes">
          <Card>
            <CardHeader>
              <CardTitle>Notes</CardTitle>
              <CardDescription>
                Investigator notes and observations
              </CardDescription>
            </CardHeader>

            <CardContent className="space-y-4">
              <Card>
                <CardContent className="pt-6">
                  Met during startup event.
                </CardContent>
              </Card>

              <Card>
                <CardContent className="pt-6">
                  Interested in healthcare and AI.
                </CardContent>
              </Card>

              <Card>
                <CardContent className="pt-6">
                  Frequently appears connected to entrepreneurship communities.
                </CardContent>
              </Card>
            </CardContent>
          </Card>
        </TabsContent>

      </Tabs>
    </div>
  );
}
