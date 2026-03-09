// src/pages/data-grid.tsx
import { useState } from "react"
import { Search } from "lucide-react"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

const transactions = [
  { id: "TX1001", customer: "Liam Johnson", status: "Paid", amount: "$250.00", date: "2026-03-01" },
  { id: "TX1002", customer: "Olivia Smith", status: "Pending", amount: "$150.00", date: "2026-03-02" },
  { id: "TX1003", customer: "Noah Williams", status: "Unpaid", amount: "$350.00", date: "2026-03-03" },
  { id: "TX1004", customer: "Emma Brown", status: "Paid", amount: "$450.00", date: "2026-03-04" },
  { id: "TX1005", customer: "James Garcia", status: "Paid", amount: "$550.00", date: "2026-03-05" },
]

export function DataGridPage() {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredData = transactions.filter((item) =>
    item.customer.toLowerCase().includes(searchTerm.toLowerCase()) ||
    item.id.toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className="flex flex-col gap-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Transactions</h1>
        <p className="text-muted-foreground">Manage and review your recent store sales.</p>
      </div>

      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <CardTitle>Recent Sales</CardTitle>
            <div className="relative w-full max-w-sm">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Search customers or IDs..."
                className="pl-8"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
              />
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[100px]">ID</TableHead>
                  <TableHead>Customer</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead className="text-right">Amount</TableHead>
                  <TableHead className="text-right">Date</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredData.length > 0 ? (
                  filteredData.map((tx) => (
                    <TableRow key={tx.id}>
                      <TableCell className="font-medium">{tx.id}</TableCell>
                      <TableCell>{tx.customer}</TableCell>
                      <TableCell>
                        <span className={`px-2 py-1 rounded-full text-xs font-semibold ${
                          tx.status === 'Paid' ? 'bg-green-100 text-green-700' : 
                          tx.status === 'Pending' ? 'bg-yellow-100 text-yellow-700' : 
                          'bg-red-100 text-red-700'
                        }`}>
                          {tx.status}
                        </span>
                      </TableCell>
                      <TableCell className="text-right">{tx.amount}</TableCell>
                      <TableCell className="text-right">{tx.date}</TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableRow>
                    <TableCell colSpan={5} className="h-24 text-center">
                      No results found.
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}