import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {ChooseDirectory, GetHenlo} from "../wailsjs/go/video/FileManager.js"
import {EventsOn, EventsOff} from "../wailsjs/runtime/runtime.js"
import {Greet} from "../wailsjs/go/main/App";

const PROGRESS_EVENT = "progress"

function App() : JSX.Element {
    const [filePath, setFilePath] = useState<string>("")
    const [outputDir, setOutputDir] = useState<string>("")
    const [isConverting, setIsConverting] = useState<boolean>(false)
    const [progress, setProgress] = useState<number>(0)

    function handleFileSelect(){
        //setFilePath("User/Sherard/Downloads/Video.mp4")
        setIsConverting(true)
        GetHenlo().then((data) => {

        })
    }

    function handleConvert(){
        ChooseDirectory().then((path) => {
            setOutputDir(path)
        })
    }

    function registerProgressEvent(){
        EventsOn(PROGRESS_EVENT, (data) => {
            setFilePath(data)
            setProgress(data)
            if (data == 100){
                setIsConverting(false)
                setFilePath("Hecho")
            }
        })
    }

    useEffect(() => {
        registerProgressEvent()
        return () => {
            EventsOff(PROGRESS_EVENT)
        }
    }, [])

    return (
        <>
            <div className="main-content">
                <h1>
                    <span>Shd</span>
                    <span>Video</span>
                    <span>Converter</span>
                </h1>
                <div className="main-body">
                    <div id="source-container" className="input-selector-container">
                        <label>Origin file</label>
                        <div className="input-group">
                            <input
                                value={filePath}
                                onChange={(e) =>
                                    setFilePath(e.target.value)}
                                id="input-source-path"
                                type="text"/>
                            <button
                                id="btn-source-choose-file"
                                disabled={isConverting}
                                onClick={handleFileSelect}>
                                Choose file
                            </button>
                        </div>
                    </div>
                    <div id="output-container" className="input-selector-container">
                        <label>Output directory</label>
                        <div className="input-group">
                            <input
                                value={outputDir}
                                onChange={(e) =>
                                    setOutputDir(e.target.value)}
                                id="input-output-path"
                                type="text"/>
                            <button
                                id="btn-convert-file"
                                onClick={handleConvert}>
                                Convert
                            </button>
                        </div>
                    </div>
                    <div className="progress-bar" style={{ width: `${progress}%` }} />
                </div>
            </div>
        </>
    )
}

export default App
