Clear-Host 
function PictureChoice {
    $PicChoice = Read-Host "Do you want to get pictures from a user or from a keyword? 1 or 2"
    if ($PicChoice -eq "1") {
        $global:PicToDL = Read-Host "What Unsplash user do you want to get pictures from?"
        $global:PicToDL = "https://source.unsplash.com/user/" + $PicToDL + "/"
    }
    else {
        $global:PicToDL = Read-Host "What keyword do you want to download pictures with. You can have multiple keywords if you separate them with commas like this: city,night"
        $global:PicToDL = "https://source.unsplash.com/random/?" + $PicToDL + "/"
        }
    
}
PictureChoice
while ($true)
{   
    Remove-Item 'c:\temp\random.jpeg'
    Invoke-WebRequest -uri $global:PicToDl -OutFile c:\temp\random.jpeg
    Function Set-WallPaper($Value)
    {

     Set-ItemProperty -path 'HKCU:\Control Panel\Desktop\' -name wallpaper -value $value
     rundll32.exe user32.dll, UpdatePerUserSystemParameters
    }

    Set-WallPaper -value c:\temp\random.jpeg

    Start-Sleep -s 60
}