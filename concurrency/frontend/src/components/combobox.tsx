"use client"

import * as React from "react"
import { Check, ChevronsUpDown } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"

const filters = [
    {
        value: "pixelate",
        label: "Pixelate"
    },
    {
        value: "grayscale",
        label: "Grayscale"
    },
    {
        value: "invert",
        label: "Invert"
    },
    {
        value: "sepia",
        label: "Sepia"
    }
]

type filterProps = {
    value: string;
    setValue: (newValue: string) => void;
}

function Combobox(props: filterProps) {
    const [open, setOpen] = React.useState(false)

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className="w-[200px] justify-between"
                >
                    {props.value
                        ? filters.find((filter) => filter.value === props.value)?.label
                        : "Select filter..."}
                    <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="w-[200px] p-0">
                <Command>
                    <CommandInput placeholder="Search filter..." />
                    <CommandEmpty>No filter found.</CommandEmpty>
                    <CommandGroup>
                        {filters.map((filter) => (
                            <CommandItem
                                key={filter.value}
                                value={filter.value}
                                onSelect={(currentValue) => {
                                    props.setValue(currentValue === props.value ? "" : currentValue)
                                    setOpen(false)
                                }}
                            >
                                <Check
                                    className={cn(
                                        "mr-2 h-4 w-4",
                                        props.value === filter.value ? "opacity-100" : "opacity-0"
                                    )}
                                />
                                {filter.label}
                            </CommandItem>
                        ))}
                    </CommandGroup>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

export default Combobox