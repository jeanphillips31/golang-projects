"use client"

import React, { useState, useEffect } from 'react';
import axios from "axios";
import {Label} from "@/components/ui/label";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {Form} from "@/components/ui/form";
import Combobox from "@/components/combobox";
import Image from "next/image";
import {decodeFromBase64} from "next/dist/build/webpack/loaders/utils";

export default function Home() {
    const [selectedFile, setSelectedFile] = useState(null)
    const [processedImageURL, setProcessedImageURL] = useState('')
    const [filterValue, setFilterValue] = React.useState("")
    
    useEffect(() => {
        // Initialize the SSE connection.
        const evtSource = new EventSource("http://localhost:8080/v1/events/imageProcessed")

        // Listen for messages to receive the processed image.
        evtSource.onmessage = (event) => {
            // Assume the event data is a URL or base64 image string.
            const dataUrl = `data:image/png;base64,${event.data}`;
            setProcessedImageURL(dataUrl)
        };

        return () => {
            evtSource.close()
        };
    }, [])
    
    
    const handleFileSelect = (event) => {
        setSelectedFile(event.target.files[0])
    };
    
    
    const handleSubmit = async(event) => {
        event.preventDefault()
        
        if (selectedFile) {
            const formData = new FormData()
            formData.append("image", selectedFile)
            formData.append("filter", filterValue)
            
            try {
                // Post the image to the /upload endpoint
                await axios.post("http://localhost:8080/v1/upload", formData).then((response) => {
                    console.log("Image uploaded Successfully")
                }).catch((err) => {
                    console.log(err)
                })
            } catch (e){
                console.error("Error:", e)
            }
        }
    }
    
    return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
          <div className={"justify-between items-center"}>
              <Combobox value={filterValue} setValue={setFilterValue}/>
              <Label htmlFor="picture">Select image</Label>
              <Input id="picture" type="file" accept={"image/png"} onChange={handleFileSelect}></Input>
              <Button onClick={handleSubmit}>Upload</Button>
          </div>
          {processedImageURL != null ? (
              <Image src={processedImageURL} alt={"processedImage"} width={500} height={500}/>
          )  : (
              <div/>
          )}
      </div>
    </main>
  );
}
